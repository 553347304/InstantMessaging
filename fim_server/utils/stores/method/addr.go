package method

import (
	"github.com/yangtizi/cz88"
	"net"
)

type addrInterface interface {
	GetAddr(string) string // 通用IP/域名   获取归属地
}

type addrServer struct {
	IP string
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Addr() addrInterface { return &addrServer{} }

func (s *addrServer) GetAddr(ip string) string {
	addr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return "nil"
	}
	_ip := addr.IP.String()
	cityAddr := cz88.GetAddress(_ip)
	return cityAddr
}
