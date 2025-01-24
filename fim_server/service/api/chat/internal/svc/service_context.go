package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/common/zero_middleware"
	"fim_server/service/api/chat/internal/config"
	"fim_server/service/api/chat/internal/middleware"
	"fim_server/service/rpc/file/file"
	"fim_server/service/rpc/file/file_rpc"
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
	FileRpc         file_rpc.FileClient
	RpcLog          service_method.ServerInterfaceLog
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DB:              src.Client().Mysql(c.System.Mysql),
		Redis:           src.Client().Redis(c.System.Redis),
		UserRpc:         client.UserClient(c.UserRpc),
		FileRpc:         file.NewFile(zrpc.MustNewClient(c.FileRpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		RpcLog:          service_method.Log(c.Name, 2),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
