package svc

import (
	"fim_server/config/core"
	"fim_server/service/rpc/group/internal/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)


type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      core.Mysql(c.System.Mysql),
		Redis:   core.Redis(c.System.Redis),
	}
}
