package ips

import (
	"fim_server/utils/stores/logs"
	"fmt"
	"net"
)

func GetAll() string {
	var ipv4 []string
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网卡信息错误", err)
		return ""
	}

	// 遍历所有网络接口
	for _, inter := range interfaces {
		addrs, errs := inter.Addrs()
		if errs != nil {
			fmt.Println("获取IP错误", errs)
			continue
		}

		// 遍历接口的地址
		for _, addr := range addrs {
			// 判断地址类型，IPv4或IPv6
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ipv4 = append(ipv4, ipNet.IP.String()+"\n")

				} else {
					// fmt.Println("IPv6:", ipNet.IP.String())
				}
			}
		}
	}
	logs.Info("IPv4:", ipv4)

	return ipv4[0]
}
