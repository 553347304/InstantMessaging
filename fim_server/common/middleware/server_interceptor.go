package middleware

import (
	"context"
	"fim_server/utils/stores/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ip := metadata.ValueFromIncomingContext(ctx, "ip")
	userId := metadata.ValueFromIncomingContext(ctx, "userId")
	ctx = context.WithValue(ctx, "ip", ip)
	ctx = context.WithValue(ctx, "user_id", userId)
	logs.Info(ip, userId)
	return handler(ctx, req)
}