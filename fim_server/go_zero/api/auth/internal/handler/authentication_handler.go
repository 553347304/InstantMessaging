package handler

import (
	"net/http"

	"fim_server/go_zero/api/auth/internal/logic"
	"fim_server/go_zero/api/auth/internal/svc"
	"fim_server/go_zero/api/auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/go_zero/server/response"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticationRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication(&req)
		response.Response(r, w, resp, err)
	}
}
