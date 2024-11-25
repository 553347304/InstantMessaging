package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	OpenLoginList []struct {
		Name string `yaml:"name"`
		Icon string `yaml:"icon"`
		Href string `yaml:"href"`
	}
	System struct {
		Mysql string
		Redis string
	}
}
