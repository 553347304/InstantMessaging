package middleware

import (
	"context"
	"fim_server/common/service/log_service"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

var interceptor = struct {
	IP     string
	UserID string
}{
	IP:     "ip",
	UserID: "user_id",
}

func ClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		interceptor.IP:     ctx.Value(interceptor.IP).(string),
		interceptor.UserID: ctx.Value(interceptor.UserID).(string),
	}))
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = context.WithValue(ctx, interceptor.IP, metadata.ValueFromIncomingContext(ctx, interceptor.IP))
	ctx = context.WithValue(ctx, interceptor.UserID, metadata.ValueFromIncomingContext(ctx, interceptor.UserID))
	return handler(ctx, req)
}

func UseMiddleware(pusher log_service.PusherServerInterface) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, interceptor.IP, httpx.GetRemoteAddr(r))
			ctx = context.WithValue(ctx, interceptor.UserID, r.Header.Get("User-Id"))
			
			writer := Writer{ResponseWriter: w}
			next(&writer, r.WithContext(ctx))
			pusher.Info(ctx, "Response", pusher.Response(w, r, writer.Body))
		}
	}
	
}
