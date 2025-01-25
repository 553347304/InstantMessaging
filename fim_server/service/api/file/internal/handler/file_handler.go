package handler

import (
	"context"
	"fim_server/models/file_models"
	"fim_server/service/api/file/internal/logic"
	"fim_server/service/api/file/internal/svc"
	"fim_server/service/api/file/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/service/server/response"
	"fim_server/utils/encryption_and_decryptio/md5s"
	"fim_server/utils/stores/files"
	"fim_server/utils/stores/logs"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"
)

func FileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		// 文件上传
		file := files.File{
			// WhiteEXT: svcCtx.Config.File.WhiteEXT,
			// BlackEXT: svcCtx.Config.File.BlackEXT,
		}.FormFile(r.FormFile("file"))

		if file.Error != "" {
			response.Response(r, w, nil, logs.Error(file.Error))
			return
		}

		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)

		// 图片已存在
		fileHash := md5s.Hash(file.Byte)
		var fileModel file_models.FileModel
		err = svcCtx.DB.Take(&fileModel, "hash = ?", fileHash).Error
		if err == nil {
			logs.Info("文件Hash重复", file.Name)
			resp.Src = fileModel.WebPath()
			response.Response(r, w, resp, err)
			return
		}

		// 用户信息
		userResponse, err := svcCtx.UserRpc.User.UserInfo(context.Background(), &user_rpc.IdList{Id: []uint32{uint32(req.UserId)}})
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		newFileModel := file_models.FileModel{
			Uid:    uuid.New(),
			UserId: req.UserId,
			Name:   file.Name,
			Size:   file.Size,
			Hash:   md5s.Hash(file.Byte),
		}

		newFileModel.Path = path.Join(svcCtx.Config.File.Path, "user",
			fmt.Sprintf("%d_%s", req.UserId, userResponse.InfoList[uint32(req.UserId)].Name),
			fmt.Sprint(newFileModel.Uid, file.Ext))

		// 写入文件
		err = files.WriteFile(newFileModel.Path, file.Byte)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		// 文件入库
		err = svcCtx.DB.Create(&newFileModel).Error
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		resp.Src = newFileModel.WebPath()
		response.Response(r, w, resp, err)
	}
}
