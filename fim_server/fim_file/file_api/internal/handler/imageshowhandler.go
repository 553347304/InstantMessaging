package handler

import (
	"net/http"
	"os"
	"path"

	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/common/response"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		filePath := path.Join("file", req.ImageType, req.ImageName)
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
