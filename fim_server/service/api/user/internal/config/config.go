package config

import (
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
	UserRpc  zrpc.RpcClientConf
	ChatRpc  zrpc.RpcClientConf
	GroupRpc zrpc.RpcClientConf
}
