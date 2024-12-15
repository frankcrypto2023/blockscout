// Process new block and unconfirmed transactions (via RPC).
package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/docopt/docopt.go"

	qngplex "qngplex/pkg"
)

func main() {
	var err error
	usage := `Process new block and unconfirmed transactions.

Usage:
  qngplex-prod [--config=<path>]
  qngplex-prod -h | --help

Options:
  -h --help     	Show this screen.
  -c <path>, --config <path>	Path to config file [default: config.json].
`

	arguments, _ := docopt.Parse(usage, nil, true, "qngplex-prod", false)

	confFile := "config.json"
	if arguments["--config"] != nil {
		confFile = arguments["--config"].(string)
	}

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %v", confFile)
	}

	conf, err := qngplex.LoadConfig(confFile)
	if err != nil {
		log.Fatalf("Can't load config file: %v", err)
	}
	pool, err := qngplex.GetRedis(conf)
	if err != nil {
		log.Fatalf("Can't connect to Redis: %v", err)
	}
	ssdb, err := qngplex.GetSSDB(conf)
	if err != nil {
		log.Fatalf("Can't connect to SSDB: %v", err)
	}
	qngplex.InitDB(conf)
	var wg sync.WaitGroup
	running := true
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	go func() {
		for sig := range cs {
			running = false
			log.Printf("Captured %v, waiting for everything to finish...\n", sig)
			wg.Wait()
			os.Exit(1)
		}
	}()

	log.Println("Catching up latest block before starting")
	timeGap := 2 * time.Millisecond

	ssdb.Get().Do("DEL", "order:latest")
	for {
		if running {
			wg.Add(1)
			done := qngplex.CatchUpLatestBlock(conf, pool, ssdb)
			wg.Done()
			if done {
				timeGap = 5 * time.Second
			}
			time.Sleep(timeGap)
		}

	}

	// Process unconfirmed transactions (power the unconfirmed txs page/API)
	// qngplex.ProcessUnconfirmedTxs(conf, pool, &running)
}
