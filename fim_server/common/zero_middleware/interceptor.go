package zero_middleware

import (
	"context"
	"fim_server/common/service/service_method"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	ip     = "ip"
	UserID = "user_id"
)

type Writer struct {
	http.ResponseWriter
	Body []byte
}

func (w *Writer) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}

func ClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		ip:     ctx.Value(ip).(string),
		UserID: ctx.Value(UserID).(string),
	}))
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}


func UseMiddlewareActionLog(l service_method.ServerInterfaceLog) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, ip, httpx.GetRemoteAddr(r))
			ctx = context.WithValue(ctx, UserID, r.Header.Get("User-ID"))
			writer := Writer{ResponseWriter: w}
			next(&writer, r.WithContext(ctx))
			l.Info(ctx, l.Response(w, r, writer.Body))
		}
	}
}

func UseMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, ip, httpx.GetRemoteAddr(r))
			ctx = context.WithValue(ctx, UserID, r.Header.Get("User-ID"))
			next(w, r.WithContext(ctx))
		}

}
