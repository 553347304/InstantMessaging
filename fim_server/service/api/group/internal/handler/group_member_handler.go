package handler

import (
	"net/http"

	"fim_server/service/api/group/internal/logic"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func GroupMemberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupMemberRequest


		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewGroupMemberLogic(r.Context(), svcCtx)
		resp, err := l.GroupMember(&req)
		response.Response(r, w, resp, err)
	}
}
