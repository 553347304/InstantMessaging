package handler

import (
	"fim_server/models/file_models"
	"fim_server/utils/stores/logs"
	"net/http"
	"os"

	"fim_server/service/api/file/internal/svc"
	"fim_server/service/api/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		var fileModel file_models.FileModel
		err := svcCtx.DB.Take(&fileModel, "uid = ?", req.Name).Error
		if err != nil {
			response.Response(r, w, nil, logs.Error("文件不存在"))
			return
		}

		file, _ := os.ReadFile(fileModel.Path)
		w.Write(file)
	}
}
