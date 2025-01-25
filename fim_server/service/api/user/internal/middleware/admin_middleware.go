package middleware

import (
	"fim_server/common/zero_middleware"
	"net/http"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		if !zero_middleware.IsAdmin(w, r) {
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
