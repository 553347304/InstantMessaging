package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"ip":      ctx.Value("ip").(string),
		"user_id": ctx.Value("user_id").(string),
	}))

	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
