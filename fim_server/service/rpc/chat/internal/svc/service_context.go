package svc

import (
	"fim_server/service/rpc/chat/internal/config"
	"fim_server/utils/src"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     src.Client().Mysql(c.System.Mysql),
	}
}
