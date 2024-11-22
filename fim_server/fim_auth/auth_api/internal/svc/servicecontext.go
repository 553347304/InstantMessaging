package svc

import (
	"fim_server/fim_auth/auth_api/internal/config"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/fim_user/user_rpc/users"
	"fim_server/global/core"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	var DB = core.Mysql(c.System.Mysql)
	var Redis = core.Redis(c.System.Redis)

	return &ServiceContext{
		Config:  c,
		DB:      DB,
		Redis:   Redis,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
