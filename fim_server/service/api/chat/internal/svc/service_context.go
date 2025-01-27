package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/common/zrpc_client"
	"fim_server/service/api/chat/internal/config"
	"fim_server/service/api/chat/internal/middleware"
	"fim_server/utils/src"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	Redis           *redis.Client
	UserRpc         zrpc_client.UserRpc
	FileRpc         zrpc_client.FileRpc
	RpcLog          service_method.ServerInterfaceLog
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              src.Client().Mysql(c.System.Mysql),
		Redis:           src.Client().Redis(c.System.Redis),
		UserRpc:         zrpc_client.Service(c.UserRpc).UserClient(),
		FileRpc:         zrpc_client.Service(c.FileRpc).FileRpc(),
		RpcLog:          service_method.Log(c.Name, 2),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
