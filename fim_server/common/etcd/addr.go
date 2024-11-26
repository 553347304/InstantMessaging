package etcd

import (
	"context"
	"fim_server/global/core"
	"fim_server/utils/ips"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

// DeliveryAddress 上送服务地址
func DeliveryAddress(etcdAddr string, serviceName string, addr string) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		logx.Error("地址错误", list)
		return
	}
	var ip = list[0]
	if ip == "0.0.0.0" {
		strings.ReplaceAll(addr, "0.0.0.0", ips.GetAll())
	}
	client := core.Etcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		logx.Error("地址上送失败", err)
		return
	}
	logx.Infof("%s:%s", serviceName, addr)
}

func GetServiceAddr(etcdAddr string, serviceName string) string {
	client := core.Etcd(etcdAddr)
	result, err := client.Get(context.Background(), serviceName)
	if err != nil || len(result.Kvs) == 0 {
		logx.Error("地址获取失败", err)
		fmt.Println(etcdAddr)
		fmt.Println(serviceName)
		return ""
	}
	return string(result.Kvs[0].Value)
}
