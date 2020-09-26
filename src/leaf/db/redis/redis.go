package redis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomodule/redigo/redis"
)

type RedisConfig struct {
	Addr string
}

var redisCfg RedisConfig

func init() {
	data, err := ioutil.ReadFile("conf/redis.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &redisCfg)
	if err != nil {
		log.Fatal("%v", err)
	}
	fmt.Println()
}

func ConnectRedis() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", redisCfg.Addr)
	return conn, err
}
