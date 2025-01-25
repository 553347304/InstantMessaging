package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	System struct {
		Mysql string
		Redis string
		Etcd  string
	}
	UserRpc        zrpc.RpcClientConf
	KqConsumerConf kq.KqConf
	KqPusherConf   struct {
		Brokers []string
		Topic   string
	}
}
