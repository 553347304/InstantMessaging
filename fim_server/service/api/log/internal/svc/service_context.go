package svc

import (
	"fim_server/config/core"
	"fim_server/service/api/log/internal/config"
	"fim_server/service/api/log/internal/middleware"
	"fim_server/service/rpc/user/user"
	"fim_server/service/rpc/user/user_rpc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         user_rpc.UserClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	KqPusherClient  *kq.Pusher
	// ActionLogs      *log_stash.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	kqClient := kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
	return &ServiceContext{
		Config:          c,
		DB:              core.Mysql(c.System.Mysql),
		UserRpc:         user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		KqPusherClient:  kqClient,
		// ActionLogs:      log_stash.NewActionPusher(kqClient, c.Name),
	}
}
