package qngplex

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"log"
	"net/http"
	"qngplex/pkg/dbmodels"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	qjson "github.com/Qitmeer/qng/core/json"
	"github.com/garyburd/redigo/redis"
)

const GenesisTx = "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b"

// Helper to make call to bitcoind RPC API
func CallQngRPC(address string, method string, id interface{}, params []interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Ignore server certificate
		},
	}

	//
	client := &http.Client{Transport: tr}
	resp, err := client.Post(address,
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalf("ReadAll: %v", err)
	//	return nil, err
	//}
	result := make(map[string]interface{})
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()
	err = decoder.Decode(&result)
	//err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}
	return result, nil
}

func GetBlockHashRPC(conf *Config, order uint) string {
	// Get the block hash
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "getBlockhash", 1, []interface{}{order})
	if err != nil {
		return ""
	}
	return res["result"].(string)
}

func GetBlockByOrderRPC(conf *Config, order uint) GetBlockOrderResult {
	// Get the block hash
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "getBlockByOrder", 1, []interface{}{order, true})
	if err != nil {
		return GetBlockOrderResult{}
	}
	b, _ := json.Marshal(res["result"])
	var r GetBlockOrderResult
	json.Unmarshal(b, &r)
	return r
}

func GetBlockCountRPC(conf *Config) uint {
	// Get the block hash
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "getBlockCount", 1, []interface{}{})
	if err != nil {
		return uint(0)
	}
	count, _ := res["result"].(json.Number).Int64()
	return uint(count)
}
func GetBlockIsBlue(conf *Config, h string) uint {
	// Get the block isblue
	// 0:not blue;  1：blue  2：Cannot confirm
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "isBlue", 1, []interface{}{h})
	if err != nil {
		return uint(0)
	}
	count, _ := res["result"].(json.Number).Int64()
	return uint(count)
}

type QngInfo struct {
	Version         int64   `json:"version"`
	ProtocolVersion int64   `json:"protocolversion"`
	Blocks          int64   `json:"blocks"`
	TimeOffset      int64   `json:"timeoffset"`
	Connections     int64   `json:"connections"`
	Proxy           string  `json:"proxy"`
	Difficulty      float64 `json:"difficulty"`
	Testnet         bool    `json:"testnet"`
	Errors          string  `json:"errors"`
}

func GetInfoRPC(conf *Config) (qngInfo *QngInfo, err error) {
	qngInfo = new(QngInfo)
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "getinfo", 1, []interface{}{})
	if err != nil {
		return
	}
	jsoninfo := res["result"].(map[string]interface{})
	qngInfo.ProtocolVersion, _ = jsoninfo["protocolversion"].(json.Number).Int64()
	qngInfo.Version, _ = jsoninfo["version"].(json.Number).Int64()
	qngInfo.Blocks, _ = jsoninfo["blocks"].(json.Number).Int64()
	qngInfo.TimeOffset, _ = jsoninfo["timeoffset"].(json.Number).Int64()
	qngInfo.Connections, _ = jsoninfo["connections"].(json.Number).Int64()
	qngInfo.Difficulty, _ = jsoninfo["difficulty"].(json.Number).Float64()
	qngInfo.Proxy = jsoninfo["proxy"].(string)
	qngInfo.Testnet = jsoninfo["testnet"].(bool)
	qngInfo.Errors = jsoninfo["errors"].(string)
	return
}

func InStrArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func ConvertBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func SaveBlockFromOrderRPC(conf *Config, pool *redis.Pool, order uint, latestOrder uint) (err error) {
	res := GetBlockByOrderRPC(conf, order)
	if res.Order != int(order) {
		return
	}
	var block dbmodels.UTXOBlock
	GetDB().Where("hash = ?", res.Hash).First(&block)
	if block.Hash != "" {
		r, err := pool.Get().Do("SET", "order:latest", strconv.Itoa(int(order)))
		fmt.Println(r, err)
		return err
	}
	block.Hash = res.Hash
	block.Height = uint(res.Height)
	block.Weight = uint(res.Weight)
	block.TransactionsCount = uint(len(res.Transactions))
	if len(res.Transactions) < 0 {
		log.Printf("Block %v has no transactions\n", res.Hash)
		return
	}
	coinbaseTx := res.Transactions[0]
	block.Coinbase = uint(coinbaseTx.Vout[0].Amount / 1e8)
	block.Confirmations = uint(res.Confirmations)
	block.Nonce = uint64(res.Pow.Nonce)
	block.Pow = res.Pow.PowName
	block.BlockOrder = order
	block.TxRoot = res.TxRoot
	block.ParentRoot = res.Parentroot
	block.StateRoot = res.StateRoot
	block.BlockTime = res.Timestamp
	block.Bits = uint64(res.Difficulty)
	block.Version = uint(res.Version)
	block.Status = dbmodels.GetBlockStatus(order, latestOrder)
	block.Txsvalid = uint(ConvertBoolToInt(res.Txsvalid))
	block.Parents = strings.Join(res.Parents, ",")
	block.Children = strings.Join(res.Children, ",")
	block.Color = dbmodels.GetBlockColor(GetBlockIsBlue(conf, res.Hash))
	err = db.Create(&block).Error
	if err != nil {
		log.Println("Error saving block:", err)
	} else {
		r, err := pool.Get().Do("SET", "order:latest", strconv.Itoa(int(order)))
		fmt.Println(r, err)
	}
	for i, tx := range res.Transactions {
		spendStatus := dbmodels.SpentStatusUnspent
		if i == 0 {
			spendStatus = int(dbmodels.GetMatureStatus(uint(conf.NodeParams.CoinbaseMaturity), order, latestOrder))
		}
		HandleTx(conf, pool, tx, uint(spendStatus))
	}
	return err
}

func HandleTx(conf *Config, pool *redis.Pool, res qjson.TxRawResult, txSpendstatus uint) error {
	isCoinbase := uint(ConvertBoolToInt(res.Vin[0].Coinbase != ""))
	preTxs := ""
	inputAmount := uint64(0)
	outputAmount := uint64(0)
	if isCoinbase != 1 {
		// handle vin
		// find pre tx , set spent status
		for _, txin := range res.Vin {
			preTxs += txin.Txid + ":" + strconv.Itoa(int(txin.Vout)) + ","
			var prexTx dbmodels.UTXOTx
			GetDB().Where("txid = ? and vout = ?", txin.Txid, txin.Vout).First(&prexTx)
			if prexTx.Txid == "" { // not exist add new
				// get raw tx
				txres, err := GetTxRPC(conf, txin.Txid)
				if err != nil {
					log.Printf("Err: %v, txid: %v not found", err, txin.Txid)
					continue
				}
				HandleTx(conf, pool, txres, dbmodels.SpentStatusSpent)
				continue
			}
			prexTx.SpentStatus = dbmodels.SpentStatusSpent
			GetDB().Save(&prexTx)
			inputAmount += txin.Value
		}
		preTxs = preTxs[:len(preTxs)-1]
	}
	for _, txout := range res.Vout {
		outputAmount += txout.Amount
	}
	// handle vout
	for vout, txout := range res.Vout {
		var tx dbmodels.UTXOTx
		GetDB().Where("txid = ? and vout = ?", res.Txid, vout).First(&tx)
		tx.Txvalid = uint(ConvertBoolToInt(res.Txsvalid))
		tx.SpentStatus = txSpendstatus
		tx.Block = res.BlockHash
		tx.BlockOrder = uint(res.BlockOrder)
		tx.Confirm = uint(res.Confirmations)
		if tx.Txid != "" { // exist
			db.Save(&tx)
			continue
		}
		txtime := time.Unix(int64(res.Time), 0)
		tx.IsCoinbase = isCoinbase
		tx.Txid = res.Txid
		tx.PreTxs = preTxs
		tx.Amount = uint(txout.Amount)
		tx.TxTime = txtime
		tx.TxSize = uint(res.Size)
		if isCoinbase != 1 {
			tx.Fee = uint(inputAmount - outputAmount)
		}

		tx.PKScript = txout.ScriptPubKey.Hex
		tx.Vout = uint(vout)
		tx.Address = txout.ScriptPubKey.Addresses[0]
		return db.Create(&tx).Error
	}
	return nil
}

// Fetch a transaction without additional info, used to fetch previous txouts when parsing txins
func GetTxOutRPC(conf *Config, tx_id string, txo_vout uint32) (txo *TxOut, err error) {
	// Hard coded genesis tx since it's not included in bitcoind RPC API
	if tx_id == GenesisTx {
		return
		//return TxData{GenesisTx, []TxIn{}, []TxOut{{"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", 5000000000}}}, nil
	}
	// Get the TX from bitcoind RPC API
	res_tx, err := CallQngRPC(conf.BitcoindRpcUrl, "getrawtransaction", 1, []interface{}{tx_id, 1})
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	txjson := res_tx["result"].(map[string]interface{})

	txojson := txjson["vout"].([]interface{})[txo_vout]
	txo = new(TxOut)
	valtmp, _ := txojson.(map[string]interface{})["value"].(json.Number).Float64()
	txo.Value = FloatToUint(valtmp)
	if txojson.(map[string]interface{})["scriptPubKey"].(map[string]interface{})["type"].(string) != "nonstandard" {
		txodata, txoisinterface := txojson.(map[string]interface{})["scriptPubKey"].(map[string]interface{})["addresses"].([]interface{})
		if txoisinterface {
			txo.Addr = txodata[0].(string)
		} else {
			txo.Addr = ""
		}
	} else {
		txo.Addr = ""
	}
	txospent := new(TxoSpent)
	txospent.Spent = false
	txo.Spent = txospent
	return
}

// Fetch a transaction via bticoind RPC API
func GetTxRPC(conf *Config, tx_id string) (tx qjson.TxRawResult, err error) {
	// Hard coded genesis tx since it's not included in bitcoind RPC API
	if tx_id == GenesisTx {
		return
		//return TxData{GenesisTx, []TxIn{}, []TxOut{{"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", 5000000000}}}, nil
	}
	// Get the TX from bitcoind RPC API
	res_tx, err := CallQngRPC(conf.BitcoindRpcUrl, "getRawTransaction", 1, []interface{}{tx_id, 1})
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	tx = res_tx["result"].(qjson.TxRawResult)

	return
}

// Fetch a transaction via bticoind RPC API
func SaveTxFromRPC(conf *Config, pool *redis.Pool, tx_id string, block *Block, tx_index int) (tx *Tx, err error) {
	c := pool.Get()
	defer c.Close()
	var wg sync.WaitGroup
	var tximut, txomut sync.Mutex
	// Hard coded genesis tx since it's not included in bitcoind RPC API
	if tx_id == GenesisTx {
		return
		//return TxData{GenesisTx, []TxIn{}, []TxOut{{"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", 5000000000}}}, nil
	}
	// Get the TX from bitcoind RPC API
	res_tx, err := CallQngRPC(conf.BitcoindRpcUrl, "getrawtransaction", 1, []interface{}{tx_id, 1})
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	txjson := res_tx["result"].(map[string]interface{})

	tx = new(Tx)
	tx.Index = uint32(tx_index)
	tx.Hash = tx_id
	tx.BlockTime = block.BlockTime
	tx.BlockHeight = block.Height
	tx.BlockHash = block.Hash
	vertmp, _ := txjson["version"].(json.Number).Int64()
	tx.Version = uint32(vertmp)
	ltimetmp, _ := txjson["locktime"].(json.Number).Int64()
	tx.LockTime = uint32(ltimetmp)
	tx.Size = uint32(len(txjson["hex"].(string)) / 2)

	total_tx_out := uint64(0)
	total_tx_in := uint64(0)

	sem := make(chan bool, 50)
	for txiindex, txijson := range txjson["vin"].([]interface{}) {
		_, coinbase := txijson.(map[string]interface{})["coinbase"]
		if !coinbase {
			wg.Add(1)
			sem <- true
			go func(pool *redis.Pool, txijson interface{}, txiindex int, total_tx_in *uint64, tx *Tx, block *Block) {
				defer wg.Done()
				defer func() { <-sem }()
				c := pool.Get()
				defer c.Close()
				txi := new(TxIn)
				txinjsonprevout := new(PrevOut)
				txinjsonprevout.Hash = txijson.(map[string]interface{})["txid"].(string)
				tmpvout, _ := txijson.(map[string]interface{})["vout"].(json.Number).Int64()
				txinjsonprevout.Vout = uint32(tmpvout)

				// Check if bitcoind is patched to fetch value/address without additional RPC call
				// cf. README
				_, bitcoindPatched := txijson.(map[string]interface{})["value"]
				if bitcoindPatched {
					pval, _ := txijson.(map[string]interface{})["value"].(json.Number).Float64()
					txinjsonprevout.Address = txijson.(map[string]interface{})["address"].(string)
					txinjsonprevout.Value = FloatToUint(pval)
				} else {
					prevout, _ := GetTxOutRPC(conf, txinjsonprevout.Hash, txinjsonprevout.Vout)

					txinjsonprevout.Address = prevout.Addr
					txinjsonprevout.Value = prevout.Value
				}
				atomic.AddUint64(total_tx_in, uint64(txinjsonprevout.Value))

				txi.PrevOut = txinjsonprevout

				tximut.Lock()
				tx.TxIns = append(tx.TxIns, txi)
				tximut.Unlock()

				txospent := new(TxoSpent)
				txospent.Spent = true
				txospent.BlockHeight = uint32(block.Height)
				txospent.InputHash = tx.Hash
				txospent.InputIndex = uint32(txiindex)

				ntxijson, _ := json.Marshal(txi)
				ntxikey := fmt.Sprintf("txi:%v:%v", tx.Hash, txiindex)

				txospentjson, _ := json.Marshal(txospent)

				c.Do("SET", ntxikey, ntxijson)
				//conn.Send("ZADD", fmt.Sprintf("txi:%v", tx.Hash), txi_index, ntxikey)

				c.Do("SET", fmt.Sprintf("txo:%v:%v:spent", txinjsonprevout.Hash, txinjsonprevout.Vout), txospentjson)

				c.Do("ZADD", fmt.Sprintf("addr:%v", txinjsonprevout.Address), block.BlockTime, tx.Hash)
				c.Do("ZADD", fmt.Sprintf("addr:%v:sent", txinjsonprevout.Address), block.BlockTime, tx.Hash)
				c.Do("HINCRBY", fmt.Sprintf("addr:%v:h", txinjsonprevout.Address), "ts", txinjsonprevout.Value)

			}(pool, txijson, txiindex, &total_tx_in, tx, block)
		}
	}
	for txo_index, txojson := range txjson["vout"].([]interface{}) {
		wg.Add(1)
		sem <- true
		go func(pool *redis.Pool, txojson interface{}, txo_index int, total_tx_out *uint64, tx *Tx, block *Block) {
			defer wg.Done()
			defer func() { <-sem }()
			c := pool.Get()
			defer c.Close()
			txo := new(TxOut)
			txoval, _ := txojson.(map[string]interface{})["value"].(json.Number).Float64()
			txo.Value = FloatToUint(txoval)
			//txo.Addr = txojson.(map[string]interface{})["scriptPubKey"].(map[string]interface{})["addresses"].([]interface{})[0].(string)

			if txojson.(map[string]interface{})["scriptPubKey"].(map[string]interface{})["type"].(string) != "nonstandard" {
				txodata, txoisinterface := txojson.(map[string]interface{})["scriptPubKey"].(map[string]interface{})["addresses"].([]interface{})
				if txoisinterface {
					txo.Addr = txodata[0].(string)
				} else {
					txo.Addr = ""
				}
			} else {
				txo.Addr = ""
			}

			txomut.Lock()
			tx.TxOuts = append(tx.TxOuts, txo)
			txomut.Unlock()
			txospent := new(TxoSpent)
			txospent.Spent = false
			txo.Spent = txospent
			//total_tx_out += uint64(txo.Value)
			atomic.AddUint64(total_tx_out, uint64(txo.Value))

			ntxojson, _ := json.Marshal(txo)
			ntxokey := fmt.Sprintf("txo:%v:%v", tx.Hash, txo_index)
			c.Do("SET", ntxokey, ntxojson)
			//conn.Send("ZADD", fmt.Sprintf("txo:%v", tx.Hash), txo_index, ntxokey)
			c.Do("ZADD", fmt.Sprintf("addr:%v", txo.Addr), block.BlockTime, tx.Hash)
			c.Do("ZADD", fmt.Sprintf("addr:%v:received", txo.Addr), block.BlockTime, tx.Hash)
			c.Do("HINCRBY", fmt.Sprintf("addr:%v:h", txo.Addr), "tr", txo.Value)
		}(pool, txojson, txo_index, &total_tx_out, tx, block)

	}

	wg.Wait()

	tx.TxOutCnt = uint32(len(tx.TxOuts))
	tx.TxInCnt = uint32(len(tx.TxIns))
	tx.TotalOut = uint64(total_tx_out)
	tx.TotalIn = uint64(total_tx_in)

	ntxjson, _ := json.Marshal(tx)
	ntxjsonkey := fmt.Sprintf("tx:%v", tx.Hash)
	c.Do("SET", ntxjsonkey, ntxjson)
	c.Do("ZADD", fmt.Sprintf("block:%v:txs", block.Hash), tx_index, ntxjsonkey)
	c.Do("ZADD", fmt.Sprintf("tx:%v:blocks", tx.Hash), tx.BlockTime, block.Hash)
	return
}

func GetRawMemPoolRPC(conf *Config) (unconfirmedtxs []string, err error) {
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "getrawmempool", 1, []interface{}{})
	if err != nil {
		return
	}
	unconfirmedtxs = []string{}
	for _, txid := range res["result"].([]interface{}) {
		unconfirmedtxs = append(unconfirmedtxs, txid.(string))
	}
	return
}

func GetRawMemPoolVerboseRPC(conf *Config) (unconfirmedtxs map[string]interface{}, err error) {
	res, err := CallQngRPC(conf.BitcoindRpcUrl, "getrawmempool", 1, []interface{}{true})
	if err != nil {
		return
	}
	unconfirmedtxs = res["result"].(map[string]interface{})
	return
}
