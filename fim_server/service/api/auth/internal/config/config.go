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
	OpenLoginList []struct {
		Name string `yaml:"name"`
		Icon string `yaml:"icon"`
		Href string `yaml:"href"`
	}
	UserRpc    zrpc.RpcClientConf
	SettingRpc zrpc.RpcClientConf
	WhiteList  []string // 白名单
}
