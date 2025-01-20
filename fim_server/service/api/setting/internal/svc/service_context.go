package svc

import (
	"fim_server/common/middleware"
	"fim_server/common/service/log_service"
	"fim_server/service/api/setting/internal/config"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	Log             log_service.PusherServerInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		Log:             log_service.NewPusher(c.Name, log_service.Action),
	}
}
