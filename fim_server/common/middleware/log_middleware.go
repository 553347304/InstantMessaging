package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := httpx.GetRemoteAddr(r)
		ctx := context.WithValue(r.Context(), "ip", ip)
		ctx = context.WithValue(ctx, "user_id", r.Header.Get("User-Id"))
		next(w, r.WithContext(ctx))
	}
}
