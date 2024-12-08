package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	System struct {
		Mysql string
		Etcd  string
	}
	UserRpc zrpc.RpcClientConf
	File    struct {
		Path     string
		MaxSize  float64
		WhiteEXT []string
		BlackEXT []string
	}
}
