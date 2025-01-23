package Admin

import (
	"net/http"

	"fim_server/service/api/group/internal/logic/Admin"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func GroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GroupList(&req)
		response.Response(r, w, resp, err)
	}
}
