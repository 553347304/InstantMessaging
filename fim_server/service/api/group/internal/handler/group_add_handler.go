package handler

import (
	"net/http"

	"fim_server/service/api/group/internal/logic"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func GroupAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewGroupAddLogic(r.Context(), svcCtx)
		resp, err := l.GroupAdd(&req)
		response.Response(r, w, resp, err)
	}
}
