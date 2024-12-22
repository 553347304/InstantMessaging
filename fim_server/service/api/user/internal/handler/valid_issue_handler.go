package handler

import (
	"net/http"

	"fim_server/service/api/user/internal/logic"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func ValidIssueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValidIssueRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewValidIssueLogic(r.Context(), svcCtx)
		resp, err := l.ValidIssue(&req)
		response.Response(r, w, resp, err)
	}
}
