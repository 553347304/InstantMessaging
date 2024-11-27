package etcd

import (
	"context"
	"fim_server/global/core"
	"fim_server/utils/stores/logs"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
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
		ip = netx.InternalIp()
		strings.ReplaceAll(addr, "0.0.0.0", ip)
	}
	client := core.Etcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		logx.Error("地址上送失败", err)
		return
	}
	logs.Info("地址上送成功   %s:%s -> %s", serviceName, addr, ip)
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
