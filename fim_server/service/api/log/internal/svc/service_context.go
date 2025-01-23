package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/service/api/log/internal/config"
	"fim_server/service/api/log/internal/middleware"
	"fim_server/service/rpc/user/client"
	"fim_server/utils/src"
	"github.com/zeromicro/go-queue/kq"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         client.UserRpc
	KqPusherClient  *kq.Pusher
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	RpcLog          service_method.ServerInterfaceLog
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              src.Client().Mysql(c.System.Mysql),
		UserRpc:         client.UserClient(c.UserRpc),
		KqPusherClient:  kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		RpcLog:          service_method.Log(c.Name, 2),
	}
}
