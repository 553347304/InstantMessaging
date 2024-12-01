package handler

import (
	"net/http"

	"fim_server/go_zero/api/setting/internal/logic"
	"fim_server/go_zero/api/setting/internal/svc"
	"fim_server/go_zero/server/response"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		response.Response(r, w, resp, err)
	}
}
