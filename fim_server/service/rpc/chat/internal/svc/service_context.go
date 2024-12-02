package svc

import (
	"fim_server/config/core"
	"fim_server/service/rpc/chat/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     core.Mysql(c.System.Mysql),
	}
}
