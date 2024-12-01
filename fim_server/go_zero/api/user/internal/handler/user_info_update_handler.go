package handler

import (
	"net/http"

	"fim_server/go_zero/api/user/internal/logic"
	"fim_server/go_zero/api/user/internal/svc"
	"fim_server/go_zero/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/go_zero/server/response"
)

func UserInfoUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewUserInfoUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoUpdate(&req)
		response.Response(r, w, resp, err)
	}
}
