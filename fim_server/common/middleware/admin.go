package middleware

import (
	"fim_server/service/server/response"
	"fim_server/utils/stores/conv"
	"net/http"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("Role")
		if role != "1" {
			response.Response(r, w, nil, conv.Type("权限验证失败").Error())
			return
		}
		next(w, r)
	}
}
