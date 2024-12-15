package qngplex

import (
	"encoding/json"
	"os"

	"github.com/Qitmeer/qng/params"
)

// Struct holding our configuration
type Config struct {
	BitcoindBlocksPath string `json:"bitcoind_blocks_path"`
	PostgresUrl        string `json:"postgres_url"`
	BitcoindRpcUrl     string `json:"bitcoind_rpc_url"`
	SsdbHost           string `json:"ssdb_host"`
	RedisHost          string `json:"redis_host"`
	LevelDbPath        string `json:"leveldb_path"`
	AppUrl             string `json:"app_url"`
	AppPort            uint   `json:"app_port"`
	AppApiRateLimited  bool   `json:"app_api_rate_limited"`
	AppTemplatesPath   string `json:"app_templates_path"`
	AppGoogleAnalytics string `json:"app_google_analytics"`
	Network            string `json:"network"`
	NodeParams         *params.Params
}

// Load configuration from json file
func LoadConfig(path string) (conf *Config, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	conf = new(Config)
	json.Unmarshal(file, conf)
	switch conf.Network {
	case "privnet":
		conf.NodeParams = &params.PrivNetParams
	case "testnet":
		conf.NodeParams = &params.TestNetParams
	case "mixnet":
		conf.NodeParams = &params.MixNetParams
	default:
		conf.NodeParams = &params.MainNetParams
	}
	return
}
