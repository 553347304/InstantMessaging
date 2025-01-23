package svc

import (
	"fim_server/service/rpc/user/internal/config"
	"fim_server/utils/src"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     src.Client().Mysql(c.System.Mysql),
		Redis:  src.Client().Redis(c.System.Redis),
	}
}
