package handler

import (
	"net/http"

	"fim_server/go_zero/api/user/internal/logic"
	"fim_server/go_zero/api/user/internal/svc"
	"fim_server/go_zero/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/go_zero/server/response"
)

func userAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAuthRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewUserAuthLogic(r.Context(), svcCtx)
		resp, err := l.UserAuth(&req)
		response.Response(r, w, resp, err)
	}
}
