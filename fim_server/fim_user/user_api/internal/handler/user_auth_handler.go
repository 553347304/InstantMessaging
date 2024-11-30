package handler

import (
	"net/http"

	"fim_server/fim_user/user_api/internal/logic"
	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/common/response"
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
