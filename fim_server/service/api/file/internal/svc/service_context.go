package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/service/api/file/internal/config"
	"fim_server/service/api/file/internal/middleware"
	"fim_server/service/rpc/user/client"
	"fim_server/utils/src"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         client.UserRpc
	RpcLog          service_method.ServerInterfaceLog
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              src.Client().Mysql(c.System.Mysql),
		UserRpc:         client.UserClient(c.UserRpc),
		RpcLog:          service_method.Log(c.Name, 2),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
