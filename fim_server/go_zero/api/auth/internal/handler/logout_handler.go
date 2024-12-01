package handler

import (
	"net/http"

	"fim_server/go_zero/api/auth/internal/logic"
	"fim_server/go_zero/api/auth/internal/svc"
	"fim_server/go_zero/server/response"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)

		token := r.Header.Get("token")

		resp, err := l.Logout(token)
		response.Response(r, w, resp, err)
	}
}
