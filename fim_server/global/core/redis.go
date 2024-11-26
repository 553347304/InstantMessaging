package core

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"strconv"
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
	db, _ := strconv.Atoi(opt[2])
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt[0],
		Password: opt[1], // no password set
		DB:       db,     // use default DB
		PoolSize: 100,    // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln("Redis连接失败: ", err.Error())
	}
	return rdb
}
