package main

import (
	"flag"
	"fmt"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	src := rest.MustNewServer(c.RestConf)
	defer src.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(src, ctx)


	fmt.Printf("Starting src at %s:%d...\n", c.Host, c.Port)
	src.Start()
}
