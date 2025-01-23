package src

import (
	"context"
	"fim_server/utils/stores/logs"
	"github.com/zeromicro/go-zero/core/netx"
	"strings"
)

type serviceInterfaceEtcd interface {
	DeliveryAddress(string, string, string)  // 上送服务地址
	GetServiceAddress(string, string) string // 获取服务地址
}
type serviceEtcd struct{}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Etcd() serviceInterfaceEtcd { return &serviceEtcd{} }

func (s *serviceEtcd) DeliveryAddress(c string, serviceName string, addr string) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		logs.Fatal("地址错误", addr)
		return
	}
	if list[0] == "0.0.0.0" {
		addr = strings.ReplaceAll(addr, "0.0.0.0", netx.InternalIp())
	}
	
	etcd := Client().Etcd(c)
	_, err := etcd.Put(context.Background(), serviceName, addr)
	if err != nil {
		logs.Fatal("地址上送失败", err)
		return
	}
	logs.InfoF("地址上送成功   %s:%s", serviceName, addr)
}
func (s *serviceEtcd) GetServiceAddress(c string, serviceName string) string {
	etcd := Client().Etcd(c)
	result, err := etcd.Get(context.Background(), serviceName)
	
	if err == nil && len(result.Kvs) > 0 {
		return string(result.Kvs[0].Value)
	}
	logs.Fatal("地址获取失败", serviceName, c, err)
	return ""
}
