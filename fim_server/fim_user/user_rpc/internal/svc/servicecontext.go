package svc

import (
	"fim_server/fim_user/user_rpc/internal/config"
	"fim_server/global/core"
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
		DB:     core.Mysql(c.System.Mysql),
		Redis:  core.Redis(c.System.Redis),
	}
}
