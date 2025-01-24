package admin

import (
	"net/http"

	"fim_server/service/api/file/internal/logic/admin"
	"fim_server/service/api/file/internal/svc"
	"fim_server/service/api/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func FileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := admin.NewFileListLogic(r.Context(), svcCtx)
		resp, err := l.FileList(&req)
		response.Response(r, w, resp, err)
	}
}
