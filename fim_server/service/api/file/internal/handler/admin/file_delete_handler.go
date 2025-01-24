package admin

import (
	"net/http"

	"fim_server/service/api/file/internal/logic/admin"
	"fim_server/service/api/file/internal/svc"
	"fim_server/service/api/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func FileDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestDelete
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := admin.NewFileDeleteLogic(r.Context(), svcCtx)
		resp, err := l.FileDelete(&req)
		response.Response(r, w, resp, err)
	}
}
