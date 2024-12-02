package main

import (
	"fim_server/utils/src/etcd"
	"flag"
	"fmt"

	"fim_server/service/api/file/internal/config"
	"fim_server/service/api/file/internal/handler"
	"fim_server/service/api/file/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/file.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	src := rest.MustNewServer(c.RestConf)
	defer src.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(src, ctx)

	etcd.DeliveryAddress(c.System.Etcd, c.Name+"_api", fmt.Sprintf("%s:%d", c.Host, c.Port))

	fmt.Printf("Starting src at %s:%d...\n", c.Host, c.Port)
	src.Start()
}
