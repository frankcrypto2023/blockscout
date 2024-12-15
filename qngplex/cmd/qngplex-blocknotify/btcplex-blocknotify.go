// Command executed by bitcoind when a new block is found,
// publish the hash over a Redis PubSub channel.
package main

import (
	"log"
	"os"

	"github.com/docopt/docopt.go"

	qngplex "qngplex/pkg"
)

func main() {
	usage := `Callback executed when bitcoind best block changes.

Usage:
  qngplex-blocknotify [--config=<path>] <hash>
  qngplex-blocknotify -h | --help

Options:
  -h --help     	Show this screen.
  -c <path>, --config <path>	Path to config file [default: config.json].
`

	arguments, _ := docopt.Parse(usage, nil, true, "qngplex-blocknotify", false)

	confFile := "config.json"
	if arguments["--config"] != nil {
		confFile = arguments["--config"].(string)
	}

	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %v", confFile)
	}

	conf, _ := qngplex.LoadConfig(confFile)
	pool, _ := qngplex.GetRedis(conf)

	conn := pool.Get()
	defer conn.Close()

	conn.Do("PUBLISH", "qngplex:blocknotify", arguments["<hash>"].(string))
}
