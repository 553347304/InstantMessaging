package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	System struct {
		Mysql string
		Redis string
		Etcd  string
	}
}

