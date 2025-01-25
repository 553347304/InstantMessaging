package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/common/zrpc_client"
	"fim_server/service/api/setting/internal/config"
	"fim_server/service/api/setting/internal/middleware"
	"fim_server/utils/src"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	RpcLog          service_method.ServerInterfaceLog
	DB              *gorm.DB
	SettingRpc      zrpc_client.SettingRpc
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	
	return &ServiceContext{
		Config:          c,
		RpcLog:          service_method.Log(c.Name, 2),
		DB:              src.Client().Mysql(c.System.Mysql),
		SettingRpc:      zrpc_client.Service(c.SettingRpc).SettingRpc(),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
