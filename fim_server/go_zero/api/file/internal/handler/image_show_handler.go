package handler

import (
	"fim_server/utils/stores/logs"
	"net/http"
	"os"
	"path"

	"fim_server/go_zero/api/file/internal/svc"
	"fim_server/go_zero/api/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/go_zero/server/response"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		filePath := path.Join("file", req.ImageType, req.ImageName)
		logs.Info(filePath)
		byteData, err := os.ReadFile(filePath)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		w.Write(byteData)

		// l := logic.NewImageShowLogic(r.Context(), svcCtx)
		// err := l.ImageShow(&req)
		// response.Response(r, w, nil, err)
	}
}
