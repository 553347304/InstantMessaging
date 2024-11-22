package core

import (
	"context"
	"fim_server/utils/convert"
	"github.com/go-redis/redis"
	"log"
	"strings"
	"time"
)

func split(s string, params []string) []string {
	var sp = strings.Split(s, " ")
	for i, i2 := range sp {
		params[i] = i2
	}
	return params
}

func Redis(c string) *redis.Client {
	opt := split(c, []string{"", "", "0", ""})
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt[0],
		Password: opt[1],              // no password set
		DB:       convert.Int(opt[2]), // use default DB
		PoolSize: 100,                 // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln("Redis连接失败: ", err.Error())
	}
	return rdb
}
