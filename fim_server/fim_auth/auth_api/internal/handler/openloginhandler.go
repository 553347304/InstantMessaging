package handler

import (
	"net/http"

	"fim_server/common/response"
	"fim_server/fim_auth/auth_api/internal/logic"
	"fim_server/fim_auth/auth_api/internal/svc"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)
		resp, err := l.Open_login()
		response.Response(r, w, resp, err)
	}
}
