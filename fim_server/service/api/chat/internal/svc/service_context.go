package svc

import (
	"fim_server/config/core"
	"fim_server/service/api/chat/internal/config"
	"fim_server/service/rpc/file/file"
	"fim_server/service/rpc/file/file_rpc"
	"fim_server/service/rpc/user/user"
	"fim_server/service/rpc/user/user_rpc"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserClient
	FileRpc file_rpc.FileClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      core.Mysql(c.System.Mysql),
		Redis:   core.Redis(c.System.Redis),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		FileRpc: file.NewFile(zrpc.MustNewClient(c.FileRpc)),
	}
}
