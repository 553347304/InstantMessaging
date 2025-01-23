package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/common/zero_middleware"
	"fim_server/service/api/user/internal/config"
	"fim_server/service/api/user/internal/middleware"
	"fim_server/service/rpc/chat/chat"
	"fim_server/service/rpc/chat/chat_rpc"
	"fim_server/service/rpc/group/group"
	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/user/client"
	"fim_server/utils/src"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	Redis           *redis.Client
	UserRpc         client.UserRpc
	ChatRpc         chat_rpc.ChatClient
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
		ChatRpc:         chat.NewChat(zrpc.MustNewClient(c.ChatRpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		GroupRpc:        group.NewGroup(zrpc.MustNewClient(c.GroupRpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		RpcLog:          service_method.Log(c.Name, 2),
	}
}
