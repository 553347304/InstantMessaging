package svc

import (
	"fim_server/common/middleware"
	"fim_server/common/service/log_service"
	"fim_server/config/core"
	"fim_server/service/api/group/internal/config"
	"fim_server/service/rpc/group/group"
	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/user/user"
	"fim_server/service/rpc/user/user_rpc"
	"net/http"
	
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	Redis           *redis.Client
	UserRpc         user_rpc.UserClient
	GroupRpc        group_rpc.GroupClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	Log             log_service.PusherServerInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              core.Mysql(c.System.Mysql),
		Redis:           core.Redis(c.System.Redis),
		UserRpc:         user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(middleware.ClientInterceptor))),
		GroupRpc:        group.NewGroup(zrpc.MustNewClient(c.GroupRpc, zrpc.WithUnaryClientInterceptor(middleware.ClientInterceptor))),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		Log:             log_service.NewPusher(c.Name, log_service.Action),
	}
}
