package config

import (
	"fim_server/utils/src"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Config struct {
	System struct {
		Mysql string `yaml:"Mysql"`
		Redis string `yaml:"Redis"`
		Etcd  string `yaml:"Etcd"`
	} `yaml:"System"`
}

var DB *gorm.DB
var Redis *redis.Client

func Init() {
	c := Yaml("settings.yaml")
	DB = src.Client().Mysql(c.System.Mysql)
	Redis = src.Client().Redis(c.System.Redis)
}
