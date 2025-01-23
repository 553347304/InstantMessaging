package Admin

import (
	"net/http"

	"fim_server/service/api/group/internal/logic/Admin"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func GroupMessageDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestDelete
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewGroupMessageDeleteLogic(r.Context(), svcCtx)
		resp, err := l.GroupMessageDelete(&req)
		response.Response(r, w, resp, err)
	}
}
