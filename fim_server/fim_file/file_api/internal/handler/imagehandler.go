package handler

import (
	"fim_server/utils/stores/logs"
	"io"
	"net/http"
	"os"
	"path"

	"fim_server/fim_file/file_api/internal/logic"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/common/response"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		file, fileHead, err := r.FormFile("image")
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		imageType := r.FormValue("image_type")
		if imageType == "" {
			response.Response(r, w, nil, logs.Error("image_type为空"))
			return
		}
		byteData, _ := io.ReadAll(file)
		fileName := fileHead.Filename
		filePath := path.Join("file", imageType, fileName)
		err = os.WriteFile(filePath, byteData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)

		resp.Url = "/" + filePath

		response.Response(r, w, resp, err)
	}
}
