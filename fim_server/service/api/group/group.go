package main

import (
	"fim_server/common/zero_middleware"
	"fim_server/utils/src"
	"flag"
	"fmt"
	
	"fim_server/service/api/group/internal/config"
	"fim_server/service/api/group/internal/handler"
	"fim_server/service/api/group/internal/svc"
	
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/group.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	
	server.Use(zero_middleware.UseMiddleware(ctx.RpcLog))
	
	src.Etcd().DeliveryAddress(c.System.Etcd, c.Name+"_api", fmt.Sprintf("%s:%d", c.Host, c.Port))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
