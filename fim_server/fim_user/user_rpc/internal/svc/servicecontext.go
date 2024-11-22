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
	var DB = core.Mysql(c.System.Mysql)
	var Redis = core.Redis(c.System.Redis)
	return &ServiceContext{
		Config: c,
		DB:     DB,
		Redis:  Redis,
	}
}
