package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/common/zero_middleware"
	"fim_server/service/api/group/internal/config"
	"fim_server/service/api/group/internal/middleware"
	"fim_server/service/rpc/group/group"
	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/user/client"
	"fim_server/utils/src"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	Redis           *redis.Client
	UserRpc         client.UserRpc
	GroupRpc        group_rpc.GroupClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	RpcLog          service_method.ServerInterfaceLog
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              src.Client().Mysql(c.System.Mysql),
		Redis:           src.Client().Redis(c.System.Redis),
		UserRpc:         client.UserClient(c.UserRpc),
		GroupRpc:        group.NewGroup(zrpc.MustNewClient(c.GroupRpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		RpcLog:          service_method.Log(c.Name, 2),
	}
}
