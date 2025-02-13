package open_api_info

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/https"
	"fmt"
	"net"
)

type addrResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Ip        string `json:"ip"`        // IP
		Country   string `json:"country"`   // 国家
		Province  string `json:"province"`  // 省
		City      string `json:"city"`      // 城市
		Districts string `json:"districts"` // 地区
		Isp       string `json:"isp"`       // 宽带
	} `json:"data"`
}

// 通过 IP/域名 获取 地址

func GetAddrByIP(ip string) string {

	addr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return ""
	}
	_ip := addr.IP.String()

	if net.ParseIP(_ip).IsPrivate() || _ip == "127.0.0.1" {
		return "内网地址"
	}

	response := https.Get("https://www.cz88.net/api/cz88/ip/base", nil, map[string]any{"ip": _ip})
	var r addrResponse
	conv.Json().Unmarshal(response.Body, &r)
	address := fmt.Sprintf("%s-%s-%s-%s %s", r.Data.Country, r.Data.Province, r.Data.City, r.Data.Districts, r.Data.Isp)
	return address
}
