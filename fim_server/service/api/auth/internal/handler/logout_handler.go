package handler

import (
	"net/http"

	"fim_server/service/api/auth/internal/logic"
	"fim_server/service/api/auth/internal/svc"
	"fim_server/service/server/response"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)

		token := r.Header.Get("token")

		resp, err := l.Logout(token)
		response.Response(r, w, resp, err)
	}
}
