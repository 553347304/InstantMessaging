package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	System struct {
		Etcd string
	}
	File struct {
		Path     string
		MaxSize  float64
		WhiteEXT []string
	}
}
