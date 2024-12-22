package handler

import (
	"net/http"

	"fim_server/service/api/user/internal/logic"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func ValidListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendValidListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewValidListLogic(r.Context(), svcCtx)
		resp, err := l.ValidList(&req)
		response.Response(r, w, resp, err)
	}
}
