package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OpenLoginList []struct {
		Name string `yaml:"name"`
		Icon string `yaml:"icon"`
		Href string `yaml:"href"`
	}
	System struct {
		Mysql string
		Redis string
	}
	UserRpc   zrpc.RpcClientConf
	Etcd      string
	WhiteList []string // 白名单
}
