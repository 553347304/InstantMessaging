package svc

import (
	"fim_server/common/service/service_method"
	"fim_server/service/api/setting/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	RpcLog          service_method.ServerInterfaceLog
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		RpcLog:          service_method.Log(c.Name, 2),
	}
}
