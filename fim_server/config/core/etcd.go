package core

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Etcd(addr string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr}, // etcd服务器地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}
