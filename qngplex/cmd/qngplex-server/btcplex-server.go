package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	qngplex "qngplex/pkg"
	"qngplex/pkg/dbmodels"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/docopt/docopt.go"
	"github.com/garyburd/redigo/redis"
	"github.com/grafov/bcast"
)

// Martini form for the search input
type searchForm struct {
	Query string `form:"q"`
}

// Struct holding page meta data, like meta tags, and some template variables
type pageMeta struct {
	Title          string
	Description    string
	Menu           string
	Block          *dbmodels.UTXOBlock
	Blocks         *[]*dbmodels.UTXOBlock
	Tx             *qngplex.Tx
	TxUnconfirmed  bool
	Txs            *[]*qngplex.Tx
	AddressData    *qngplex.AddressData
	LastHeight     uint
	CurrentHeight  uint
	Error          string
	Price          float64
	PaginationData *PaginationData
	Analytics      string
	QngplexSynced  bool
	QngInfo        *qngplex.QngInfo
}

type PaginationData struct {
	CurrentPage int
	MaxPage     int
	Next        int
	Prev        int
	Pages       []struct{}
}

type RedisWrapper struct {
	Pool *redis.Pool
}

const (
	ratelimitwindow = 3600
	ratelimitcnt    = 3600
	txperpage       = 20
	synctimeout     = 60 * 8
)

var conf *qngplex.Config

// Keep track of the number of active SSE client
var activeclientsmutex sync.Mutex
var activeclients uint

func incrementClient() {
	activeclientsmutex.Lock()
	defer activeclientsmutex.Unlock()

	activeclients++
}

func decrementClient() {
	activeclientsmutex.Lock()
	defer activeclientsmutex.Unlock()

	activeclients--
}

// Used to rate-limit the API
func rateLimited(rediswrapper *RedisWrapper, ip string) (bool, int, int) {
	conn := rediswrapper.Pool.Get()
	defer conn.Close()
	reset := int(time.Now().UTC().Unix()/ratelimitwindow*ratelimitwindow + ratelimitwindow)
	ipkey := fmt.Sprintf("rl:%v:%v", ip, reset)
	cnt, _ := redis.Int(conn.Do("GET", ipkey))
	if cnt > ratelimitcnt {
		return true, cnt, reset
	} else {
		conn.Send("MULTI")
		conn.Send("INCR", ipkey)
		conn.Send("EXPIREAT", ipkey, reset+ratelimitwindow)
		conn.Do("EXEC")
		cnt += 1
		return false, cnt, reset
	}
}

func bcastToRedisPubSub(pool *redis.Pool, psgroup *bcast.Group, redischannel string) {
	conn := pool.Get()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(redischannel)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			h1 := psgroup.Join()
			h1.Send(string(v.Data))
			h1.Close()
		}
	}
}

func addHATEOAS(links map[string]map[string]string, key string, link string) map[string]map[string]string {
	newlink := map[string]string{}
	newlink["href"] = link
	links[key] = newlink
	return links
}

func initHATEOAS(links map[string]map[string]string, req *http.Request) map[string]map[string]string {
	links = map[string]map[string]string{}
	return addHATEOAS(links, "self", fmt.Sprintf("%v%v", conf.AppUrl, req.URL.String()))
}

func N(n int) []struct{} {
	return make([]struct{}, n)
}

func main() {
	var err error
	var latestheight, latestheightcache int
	var blockscached *[]*dbmodels.UTXOBlock
	usage := `qngplex webapp/API server.

Usage:
  qngplex-server [--config=<path>]
  qngplex-server -h | --help

Options:
  -h --help         Show this screen.
  -c <path>, --config <path>    Path to config file [default: config.json].
`
	arguments, _ := docopt.Parse(usage, nil, true, "qngplex-server", false)

	confFile := "config.json"
	if arguments["--config"] != nil {
		confFile = arguments["--config"].(string)
	}

	log.Println("Starting qngplex-server")

	conf, err = qngplex.LoadConfig(confFile)
	if err != nil {
		log.Fatalf("Can't load config file: %v\n", err)
	}
	qngplex.InitDB(conf)
	// Used for pub/sub in the webapp and data like latest processed height
	pool, err := qngplex.GetRedis(conf)
	if err != nil {
		log.Fatalf("Can't connect to Redis: %v\n", err)
	}

	// Due to args injection I can't use two *redis.Pool with maritini
	rediswrapper := new(RedisWrapper)
	rediswrapper.Pool = pool

	ssdb, err := qngplex.GetSSDB(conf)
	if err != nil {
		log.Fatalf("Can't connect to SSDB: %v\n", err)
	}

	// Setup some pubsub:

	// Compute the unconfirmed transaction count in a ticker
	utxscnt := 0
	utxscntticker := time.NewTicker(1 * time.Second)
	go func(pool *redis.Pool, utxscnt *int) {
		c := pool.Get()
		defer c.Close()
		for _ = range utxscntticker.C {
			*utxscnt, _ = redis.Int(c.Do("ZCARD", "qngplex:rawmempool"))
		}
	}(pool, &utxscnt)

	// Pool the latest height from qngplex db,
	// also track the status/check if qngplex goes out of sync
	latestheightticker := time.NewTicker(1 * time.Second)
	checkinprogress := false
	qngorder := qngplex.GetBlockCountRPC(conf)
	qngplexsynced := true
	go func(pool *redis.Pool, latestheight *int) {
		c := pool.Get()
		defer c.Close()
		for _ = range latestheightticker.C {
			*latestheight, _ = redis.Int(c.Do("GET", "order:latest"))

			if latestheightcache != *latestheight {
				blocks, _ := qngplex.GetLastXBlocks(ssdb, uint(*latestheight), uint(*latestheight-30))
				log.Println("Re-building homepage blocks cache", latestheight, len(blocks))
				blockscached = &blocks
				latestheightcache = *latestheight
			}

			qngorder = qngplex.GetBlockCountRPC(conf)
			if uint(latestheightcache) != qngorder && !checkinprogress && qngplexsynced {
				checkinprogress = true
				go func(checkinprogress *bool) {
					if qngorder-uint(latestheightcache) > 20 {
						qngplexsynced = false
						log.Printf("CRITICAL: OUT OF SYNC / qngplex:%v, qng:%v\n", latestheightcache, qngorder)
					} else {
						log.Println("WARNING: qngplex Out of sync, waiting before another check")
						time.Sleep(synctimeout * time.Second)
						if qngplexsynced && uint(latestheightcache) != qngorder {
							qngplexsynced = false
							log.Printf("CRITICAL: OUT OF SYNC / qngplex:%v, qng:%v\n", latestheightcache, qngorder)
						}
					}
					*checkinprogress = false
				}(&checkinprogress)
			}
			if uint(latestheightcache) == qngorder && !qngplexsynced {
				log.Println("INFO: Sync with qng done")
				qngplexsynced = true
			}
		}
	}(ssdb, &latestheight)

	// PubSub channel for blocknotify qng RPC like
	blocknotifygroup := bcast.NewGroup()
	go blocknotifygroup.Broadcast(0)
	go bcastToRedisPubSub(pool, blocknotifygroup, "qngplex:blocknotify2")

	// PubSub channel for unconfirmed txs / rawmemorypool
	utxgroup := bcast.NewGroup()
	go utxgroup.Broadcast(0)
	go bcastToRedisPubSub(pool, utxgroup, "qngplex:utxs")
	// TODO Ticker for utxs count => events_unconfirmed

	newblockgroup := bcast.NewGroup()
	go newblockgroup.Broadcast(0)
	go bcastToRedisPubSub(pool, newblockgroup, "qngplex:newblock")

	qngplexsyncedgroup := bcast.NewGroup()
	go qngplexsyncedgroup.Broadcast(0)

	// Go template helper
	appHelpers := template.FuncMap{
		"cut": func(addr string, length int) string {
			return fmt.Sprintf("%v...", addr[:length])
		},
		"cutmiddle": func(addr string, length int) string {
			return fmt.Sprintf("%v...%v", addr[:length], addr[len(addr)-length:])
		},
		"tokb": func(size uint32) string {
			return fmt.Sprintf("%.3f", float32(size)/1024)
		},
		"computefee": func(tx *qngplex.Tx) string {
			if tx.TotalIn == 0 {
				return "0"
			}
			return fmt.Sprintf("%v", float32(tx.TotalIn-tx.TotalOut)/1e8)
		},
		"generationmsg": func(tx *qngplex.Tx) string {
			reward := tx.TxOuts[0].Value
			// fee := float64(tx.TotalOut-uint64(reward)) / 1e8
			return fmt.Sprintf("%v MEER ", reward)
		},
		"tobtc": func(val uint64) string {
			return fmt.Sprintf("%.8f", float64(val)/1e8)
		},
		"showBlockStatus": func(status uint) string {
			switch status {
			case dbmodels.BLOCK_STATUS_STABLE:
				return "stable"
			case dbmodels.BLOCK_STATUS_UNSTABLE:
				return "unstable"
			default:
				return "unknown"
			}
		},
		"inttobtc": func(val int64) string {
			return fmt.Sprintf("%.8f", float64(val)/1e8)
		},
		"formatprevout": func(prevout *qngplex.PrevOut) string {
			return fmt.Sprintf("%v:%v", prevout.Hash, prevout.Vout)
		},
		"formattime": func(ts uint32) string {
			return fmt.Sprintf("%v", time.Unix(int64(ts), 0).UTC())
		},
		"formatiso": func(ts uint32) string {
			return fmt.Sprintf("%v", time.Unix(int64(ts), 0).Format(time.RFC3339))
		},
		"sub": func(h, p uint) uint {
			return h - p
		},
		"add": func(h, p uint) uint {
			return h + p
		},
		"iadd": func(h, p int) int {
			return h + p
		},
		"toarr": func(s string) []string {
			return strings.Split(s, ",")
		},
		"confirmation": func(hash string, height uint) uint {
			bm, _ := qngplex.NewBlockMeta(ssdb, hash)
			if bm.Main == false {
				return 0
			}
			return uint(latestheight) - height + 1
		},
		"is_orphaned": func(block *qngplex.Block) bool {
			if block.Height == uint(latestheight) {
				return false
			}
			return !block.Main
		},
	}

	m := martini.Classic()
	m.Map(rediswrapper)
	m.Map(ssdb)

	tmpldir := "templates"
	if conf.AppTemplatesPath != "" {
		tmpldir = conf.AppTemplatesPath
	}
	m.Use(render.Renderer(render.Options{
		Directory: tmpldir,
		Layout:    "layout",
		Funcs:     []template.FuncMap{appHelpers},
	}))

	// We rate limit the API if enabled in the config
	if conf.AppApiRateLimited {
		m.Use(func(res http.ResponseWriter, req *http.Request, rediswrapper *RedisWrapper, log *log.Logger) {
			remoteIP := strings.Split(req.RemoteAddr, ":")[0]
			_, xforwardedfor := req.Header["X-Forwarded-For"]
			if xforwardedfor {
				remoteIP = req.Header["X-Forwarded-For"][1]
			}
			log.Printf("R:%v\nip:%+v\n", time.Now(), remoteIP)
			if strings.Contains(req.RequestURI, "/api/") {
				ratelimited, cnt, reset := rateLimited(rediswrapper, remoteIP)
				// Set X-RateLimit-* Header
				res.Header().Set("X-RateLimit-Limit", strconv.Itoa(ratelimitcnt))
				res.Header().Set("X-RateLimit-Remaining", strconv.Itoa(ratelimitcnt-cnt))
				res.Header().Set("X-RateLimit-Reset", strconv.Itoa(reset))
				// Set CORS header
				res.Header().Set("Access-Control-Expose-Headers", " X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset")
				res.Header().Set("Access-Control-Allow-Origin", "*")

				if ratelimited {
					res.WriteHeader(429)
				}
			}
		})
	}

	// Don't want Google to crawl API
	m.Get("/robots.txt", func() string {
		return "User-agent: *\nDisallow: /api"
	})

	m.Get("/", func(r render.Render, db *redis.Pool) {
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.Blocks = blockscached
		pm.Title = "Latest Qng blocks"
		pm.Description = "Qng block chain explorer with JSON API"
		pm.Menu = "latest_blocks"
		pm.LastHeight = uint(latestheight)
		pm.Analytics = conf.AppGoogleAnalytics
		r.HTML(200, "index", &pm)
	})

	m.Get("/blocks/:currentorder", func(params martini.Params, r render.Render, db *redis.Pool) {
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		currentorder, _ := strconv.ParseUint(params["currentorder"], 10, 0)
		blocks, _ := qngplex.GetLastXBlocks(db, uint(currentorder), uint(currentorder-30))
		pm.Blocks = &blocks
		pm.Title = "Qng blocks"
		pm.Menu = "blocks"
		pm.LastHeight = uint(latestheight)
		pm.CurrentHeight = uint(currentorder)
		pm.Analytics = conf.AppGoogleAnalytics
		r.HTML(200, "blocks", &pm)
	})

	m.Get("/block/:hash", func(params martini.Params, r render.Render, db *redis.Pool) {
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.LastHeight = uint(latestheight)
		block, _ := qngplex.GetBlockCachedByHash(db, params["hash"])
		pm.Block = block
		pm.Title = fmt.Sprintf("Qng block #%v", block.Height)
		pm.Description = fmt.Sprintf("Qng block #%v summary and related transactions", block.Height)
		pm.Analytics = conf.AppGoogleAnalytics
		txs, _ := qngplex.GetTxsByBlock(db, params["hash"])
		pm.Txs = &txs
		r.HTML(200, "block", &pm)
	})

	m.Get("/api/block/:hash", func(params martini.Params, r render.Render, db *redis.Pool, req *http.Request) {
		block, _ := qngplex.GetBlockCachedByHash(db, params["hash"])
		// block.Links = initHATEOAS(block.Links, req)
		// if block.Parent != "" {
		// 	block.Links = addHATEOAS(block.Links, "previous_block", fmt.Sprintf("%v/api/block/%v", conf.AppUrl, block.Parent))
		// }
		// if block.Next != "" {
		// 	block.Links = addHATEOAS(block.Links, "next_block", fmt.Sprintf("%v/api/block/%v", conf.AppUrl, block.Next))
		// }
		r.JSON(200, block)
	})

	m.Get("/unconfirmed-transactions", func(params martini.Params, r render.Render, db *redis.Pool, rdb *RedisWrapper) {
		//rpool := rdb.Pool
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.LastHeight = uint(latestheight)
		pm.Menu = "utxs"
		pm.Title = "Unconfirmed transactions"
		pm.Description = "Transactions waiting to be included in a Qng block, updated in real time."
		//utxs, _ := qngplex.GetUnconfirmedTxs(rpool)
		pm.Txs = &[]*qngplex.Tx{}
		pm.Analytics = conf.AppGoogleAnalytics
		r.HTML(200, "unconfirmed-transactions", &pm)
	})

	m.Get("/tx/:hash", func(params martini.Params, r render.Render, db *redis.Pool, rdb *RedisWrapper) {
		var tx *qngplex.Tx
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.LastHeight = uint(latestheight)

		tx, _ = qngplex.GetTx(db, params["hash"])
		pm.Tx = tx
		pm.TxUnconfirmed = tx.BlockHeight == 0
		pm.Title = fmt.Sprintf("Qng transaction %v", tx.Hash)
		pm.Description = fmt.Sprintf("Qng transaction %v summary.", tx.Hash)
		pm.Analytics = conf.AppGoogleAnalytics
		r.HTML(200, "tx", pm)
	})
	m.Get("/api/tx/:hash", func(params martini.Params, r render.Render, db *redis.Pool, rdb *RedisWrapper, req *http.Request) {
		var tx *qngplex.Tx
		tx, _ = qngplex.GetTx(db, params["hash"])
		tx.Links = initHATEOAS(tx.Links, req)
		if tx.BlockHash != "" {
			tx.Links = addHATEOAS(tx.Links, "block", fmt.Sprintf("%v/api/block/%v", conf.AppUrl, tx.BlockHash))
		}
		r.JSON(200, tx)
	})

	m.Get("/address/:address", func(params martini.Params, r render.Render, db *redis.Pool, req *http.Request) {
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.LastHeight = uint(latestheight)
		pm.PaginationData = new(PaginationData)
		pm.Title = fmt.Sprintf("Qng address %v", params["address"])
		pm.Description = fmt.Sprintf("Transactions and summary for the Qng address %v.", params["address"])
		// AddressData
		addressdata, _ := qngplex.GetAddress(db, params["address"])
		pm.AddressData = addressdata
		// Pagination
		d := float64(addressdata.TxCnt) / float64(txperpage)
		pm.PaginationData.MaxPage = int(math.Ceil(d))
		currentPage := req.URL.Query().Get("page")
		if currentPage == "" {
			currentPage = "1"
		}
		pm.PaginationData.CurrentPage, _ = strconv.Atoi(currentPage)
		pm.PaginationData.Pages = N(pm.PaginationData.MaxPage)
		pm.PaginationData.Next = 0
		pm.PaginationData.Prev = 0
		if pm.PaginationData.CurrentPage > 1 {
			pm.PaginationData.Prev = pm.PaginationData.CurrentPage - 1
		}
		if pm.PaginationData.CurrentPage < pm.PaginationData.MaxPage {
			pm.PaginationData.Next = pm.PaginationData.CurrentPage + 1
		}
		fmt.Printf("%+v\n", pm.PaginationData)
		// Fetch txs given the pagination
		addressdata.FetchTxs(db, txperpage*(pm.PaginationData.CurrentPage-1), txperpage*pm.PaginationData.CurrentPage)
		r.HTML(200, "address", pm)
	})
	m.Get("/api/address/:address", func(params martini.Params, r render.Render, db *redis.Pool, req *http.Request) {
		addressdata, _ := qngplex.GetAddress(db, params["address"])
		lastPage := int(math.Ceil(float64(addressdata.TxCnt) / float64(txperpage)))
		currentPageStr := req.URL.Query().Get("page")
		if currentPageStr == "" {
			currentPageStr = "1"
		}
		currentPage, _ := strconv.Atoi(currentPageStr)
		// HATEOS section
		addressdata.Links = initHATEOAS(addressdata.Links, req)
		pageurl := "%v/api/address/%v?page=%v"
		if currentPage < lastPage {
			addressdata.Links = addHATEOAS(addressdata.Links, "last", fmt.Sprintf(pageurl, conf.AppUrl, params["address"], lastPage))
			addressdata.Links = addHATEOAS(addressdata.Links, "next", fmt.Sprintf(pageurl, conf.AppUrl, params["address"], currentPage+1))
		}
		if currentPage > 1 {
			addressdata.Links = addHATEOAS(addressdata.Links, "previous", fmt.Sprintf(pageurl, conf.AppUrl, params["address"], currentPage-1))
		}
		addressdata.FetchTxs(db, txperpage*(currentPage-1), txperpage*currentPage)
		r.JSON(200, addressdata)
	})

	m.Get("/about", func(r render.Render) {
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.LastHeight = uint(latestheight)
		pm.Title = "About"
		pm.Description = "Learn more about qngplex, an open source Qng blockchain browser written in Go."
		pm.Menu = "about"
		pm.Analytics = conf.AppGoogleAnalytics
		r.HTML(200, "about", pm)
	})

	m.Get("/status", func(r render.Render) {
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		pm.LastHeight = uint(latestheight)
		pm.Title = "Status"
		pm.Description = "qngplex status page."
		pm.Menu = "status"
		pm.Analytics = conf.AppGoogleAnalytics
		qngplexinfo, _ := qngplex.GetInfoRPC(conf)
		pm.QngInfo = qngplexinfo
		r.HTML(200, "status", pm)
	})

	m.Post("/search", binding.Form(searchForm{}), binding.ErrorHandler, func(search searchForm, r render.Render, db *redis.Pool, rdb *RedisWrapper) {
		rpool := rdb.Pool
		pm := new(pageMeta)
		pm.QngplexSynced = qngplexsynced
		// Check if the query isa block height
		isblockheight, hash := qngplex.IsBlockHeight(db, search.Query)
		if isblockheight && hash != "" {
			r.Redirect(fmt.Sprintf("/block/%v", hash))
		}
		// Check if the query is block hash
		isblockhash, hash := qngplex.IsBlockHash(db, search.Query)
		if isblockhash {
			r.Redirect(fmt.Sprintf("/block/%v", hash))
		}
		// Check for TX
		istxhash, txhash := qngplex.IsTxHash(db, search.Query)
		if istxhash {
			r.Redirect(fmt.Sprintf("/tx/%v", txhash))
		}
		isutx, txhash := qngplex.IsUnconfirmedTx(rpool, search.Query)
		if isutx {
			r.Redirect(fmt.Sprintf("/tx/%v", txhash))
		}
		// Check for Qng address
		isaddress, address := qngplex.IsAddress(search.Query)
		if isaddress {
			r.Redirect(fmt.Sprintf("/address/%v", address))
		}
		pm.Title = "Search"
		pm.Error = "Nothing found"
		pm.Analytics = conf.AppGoogleAnalytics
		r.HTML(200, "search", pm)
	})

	m.Get("/api/getblockcount", func(r render.Render) {
		r.JSON(200, latestheight)
	})

	//	m.Get("/api/latesthash", func(r render.Render) {
	//		r.JSON(200, latesthash)
	//	})

	m.Get("/api/getblockhash/:height", func(r render.Render, params martini.Params, db *redis.Pool) {
		height, _ := strconv.ParseUint(params["height"], 10, 0)
		blockhash, _ := qngplex.GetBlockHash(db, uint(height))
		r.JSON(200, blockhash)
	})

	m.Get("/api/getreceivedbyaddress/:address", func(r render.Render, params martini.Params, db *redis.Pool) {
		res, _ := qngplex.GetReceivedByAddress(db, params["address"])
		r.JSON(200, res)
	})

	m.Get("/api/getsentbyaddress/:address", func(r render.Render, params martini.Params, db *redis.Pool) {
		res, _ := qngplex.GetSentByAddress(db, params["address"])
		r.JSON(200, res)
	})

	m.Get("/api/addressbalance/:address", func(r render.Render, params martini.Params, db *redis.Pool) {
		res, _ := qngplex.AddressBalance(db, params["address"])
		r.JSON(200, res)
	})

	m.Get("/api/checkaddress/:address", func(params martini.Params, r render.Render) {
		valid, _ := qngplex.ValidA58([]byte(params["address"]))
		r.JSON(200, valid)
	})

	m.Get("/api/blocknotify", func(w http.ResponseWriter, r *http.Request) {
		incrementClient()
		defer decrementClient()
		running := true
		notifier := w.(http.CloseNotifier).CloseNotify()
		timer := time.NewTimer(time.Second * 1800)

		f, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		bnotifier := blocknotifygroup.Join()
		defer bnotifier.Close()

		var ls interface{}
		for {
			if running {
				select {
				case ls = <-bnotifier.Read:
					io.WriteString(w, fmt.Sprintf("data: %v\n\n", ls.(string)))
					f.Flush()
				case <-notifier:
					running = false
					log.Println("CLOSED")
					break
				case <-timer.C:
					running = false
					log.Println("TimeOUT")
				}
			} else {
				log.Println("DONE")
				break
			}
		}
	})

	m.Get("/api/utxs/:address", func(w http.ResponseWriter, params martini.Params, r *http.Request, rdb *RedisWrapper) {
		incrementClient()
		defer decrementClient()
		rpool := rdb.Pool
		running := true
		notifier := w.(http.CloseNotifier).CloseNotify()
		timer := time.NewTimer(time.Second * 3600)

		f, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		utxs := make(chan string)
		go func(rpool *redis.Pool, utxs chan<- string) {
			conn := rpool.Get()
			defer conn.Close()
			psc := redis.PubSubConn{Conn: conn}
			psc.Subscribe(fmt.Sprintf("addr:%v:txs", params["address"]))
			for {
				switch v := psc.Receive().(type) {
				case redis.Message:
					utxs <- string(v.Data)
				}
			}
		}(rpool, utxs)

		var ls string
		for {
			if running {
				select {
				case ls = <-utxs:
					io.WriteString(w, fmt.Sprintf("data: %v\n\n", ls))
					f.Flush()
				case <-notifier:
					running = false
					log.Println("CLOSED")
					break
				case <-timer.C:
					running = false
					log.Println("TimeOUT")
				}
			} else {
				log.Println("DONE")
				break
			}
		}
	})

	m.Get("/api/utxs", func(w http.ResponseWriter, r *http.Request) {
		incrementClient()
		defer decrementClient()
		running := true
		notifier := w.(http.CloseNotifier).CloseNotify()
		timer := time.NewTimer(time.Second * 3600)

		f, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		utx := utxgroup.Join()
		defer utx.Close()

		var ls interface{}
		for {
			if running {
				select {
				case ls = <-utx.Read:
					io.WriteString(w, fmt.Sprintf("data: %v\n\n", ls.(string)))
					f.Flush()
				case <-notifier:
					running = false
					log.Println("CLOSED")
					break
				case <-timer.C:
					running = false
					log.Println("TimeOUT")
				}
			} else {
				log.Println("DONE")
				break
			}
		}
	})

	m.Get("/events", func(w http.ResponseWriter, r *http.Request) {
		running := true
		notifier := w.(http.CloseNotifier).CloseNotify()
		timer := time.NewTimer(time.Second * 8400)

		f, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		newblockg := newblockgroup.Join()
		defer newblockg.Close()
		var ls interface{}
		for {
			if running {
				select {
				case ls = <-newblockg.Read:
					io.WriteString(w, fmt.Sprintf("data: %v\n\n", ls.(string)))
					f.Flush()
				case <-notifier:
					running = false
					log.Println("CLOSED")
					break
				case <-timer.C:
					running = false
					log.Println("TimeOUT")
				}
			} else {
				log.Println("DONE")
				break
			}
		}
	})

	m.Get("/events_unconfirmed", func(w http.ResponseWriter, r *http.Request) {
		running := true
		notifier := w.(http.CloseNotifier).CloseNotify()
		timer := time.NewTimer(time.Second * 3600)

		f, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		utx := utxgroup.Join()
		defer utx.Close()

		var ls interface{}
		for {
			if running {
				select {
				case ls = <-utx.Read:
					buf := bytes.NewBufferString("")
					utx := new(qngplex.Tx)
					json.Unmarshal([]byte(ls.(string)), utx)
					t := template.New("").Funcs(appHelpers)
					utxtmpl, _ := ioutil.ReadFile(fmt.Sprintf("%v/utx.tmpl", tmpldir))
					t, err := t.Parse(string(utxtmpl))
					if err != nil {
						log.Printf("ERR:%v", err)
					}

					err = t.Execute(buf, utx)
					if err != nil {
						log.Printf("ERR EXEC:%v", err)
					}
					res := map[string]interface{}{}
					// Full unconfirmed cnt from global variables
					res["cnt"] = utxscnt
					// HTML template of the transaction
					res["tmpl"] = buf.String()
					// Last updated time
					res["time"] = time.Now().UTC().Format(time.RFC3339)
					resjson, _ := json.Marshal(res)
					io.WriteString(w, fmt.Sprintf("data: %v\n\n", string(resjson)))
					f.Flush()
				case <-notifier:
					running = false
					log.Println("CLOSED")
					break
				case <-timer.C:
					running = false
					log.Println("TimeOUT")
				}
			} else {
				log.Println("DONE")
				break
			}
		}
	})

	m.Get("/api/info", func(r render.Render) {
		activeclientsmutex.Lock()
		defer activeclientsmutex.Unlock()
		qngplexinfo, _ := qngplex.GetInfoRPC(conf)
		r.JSON(200, map[string]interface{}{"activeclients": activeclients, "info": qngplexinfo})
	})

	log.Printf("Listening on port: %v\n", conf.AppPort)
	http.ListenAndServe(fmt.Sprintf(":%v", conf.AppPort), m)
}
