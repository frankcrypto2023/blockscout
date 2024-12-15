package qngplex

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func CatchUpLatestBlock(conf *Config, rpool *redis.Pool, spool *redis.Pool) (done bool) {
	blockcount := GetBlockCountRPC(conf)
	sc := spool.Get()
	defer sc.Close()
	latestorder, _ := redis.Int(sc.Do("GET", "order:latest"))
	if uint(latestorder) == blockcount-1 {
		return true
	}
	log.Printf("Catch up block: %v\n", latestorder+1)
	SaveBlockFromOrderRPC(conf, spool, uint(latestorder+1), uint(latestorder))
	return false
}

// func ProcessNewBlock(conf *Config, rpool *redis.Pool, spool *redis.Pool) {
// 	log.Println("ProcessNewBlock startup")
// 	conn := rpool.Get()
// 	defer conn.Close()
// 	psc := redis.PubSubConn{Conn: conn}
// 	psc.Subscribe("qngplex:blocknotify")
// 	for {
// 		switch v := psc.Receive().(type) {
// 		case redis.Message:
// 			hash := string(v.Data)
// 			log.Printf("Processing new block: %v\n", hash)
// 			c := rpool.Get()
// 			newblock, err := SaveBlockFromRPC(conf, spool, hash)
// 			if err != nil {
// 				log.Printf("Error processing new block: %v\n", err)
// 			} else {
// 				// Once the block is processed, we can publish it as qngplex own blocknotify
// 				c.Do("PUBLISH", "qngplex:blocknotify2", hash)
// 				newblockjson, _ := json.Marshal(newblock)
// 				c.Do("PUBLISH", "qngplex:newblock", string(newblockjson))
// 			}
// 			c.Close()
// 		}
// 	}
// }
