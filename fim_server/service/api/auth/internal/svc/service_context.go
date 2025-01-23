package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/service/api/auth/internal/config"
	"fim_server/service/rpc/user/client"
	"fim_server/utils/src"
	
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	Redis           *redis.Client
	UserRpc         client.UserRpc
	RpcLog          service_method.ServerInterfaceLog
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              src.Client().Mysql(c.System.Mysql),
		Redis:           src.Client().Redis(c.System.Redis),
		UserRpc:         client.UserClient(c.UserRpc),
		RpcLog:          service_method.Log(c.Name, 1),
	}
}
