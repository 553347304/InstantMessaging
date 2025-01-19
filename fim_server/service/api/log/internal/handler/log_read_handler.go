package handler

import (
	"net/http"

	"fim_server/service/api/log/internal/logic"
	"fim_server/service/api/log/internal/svc"
	"fim_server/service/api/log/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func LogReadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogReadRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewLogReadLogic(r.Context(), svcCtx)
		resp, err := l.LogRead(&req)
		response.Response(r, w, resp, err)
	}
}
