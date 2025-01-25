package main

import (
	"fim_server/common/zero_middleware"
	"fim_server/service/api/log/internal/config"
	"fim_server/service/api/log/internal/handler"
	"fim_server/service/api/log/internal/svc"
	"fim_server/utils/src"

	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/log.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 设置全局中间件
	server.Use(zero_middleware.UseMiddleware(ctx.RpcLog))

	// kq service
	// serviceGroup := service.NewServiceGroup()
	// defer serviceGroup.Stop()
	// for _, mq := range mqs.Consumers(c, context.Background(), ctx) {
	// 	serviceGroup.Add(mq)
	// }
	// go serviceGroup.Start()

	src.Etcd().DeliveryAddress(c.System.Etcd, c.Name+"_api", fmt.Sprintf("%s:%d", c.Host, c.Port))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
