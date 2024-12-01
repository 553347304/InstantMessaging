package main

import (
	"fim_server/go_zero/rpc/user/internal/config"
	"fim_server/go_zero/rpc/user/internal/server"
	"fim_server/go_zero/rpc/user/internal/svc"
	"fim_server/go_zero/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_rpc.RegisterUsersServer(grpcServer, server.NewUsersServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logs.Info("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
