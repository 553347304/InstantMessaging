package admin

import (
	"context"
	"fim_server/models/file_models"
	"fim_server/utils/stores/method"

	"fim_server/service/api/file/internal/svc"
	"fim_server/service/api/file/internal/types"
)

type FileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDeleteLogic) FileDelete(req *types.RequestDelete) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	var fileList []file_models.FileModel
	l.svcCtx.DB.Find(&fileList, "id in ?", req.IdList).Delete(&fileList)

	for _, model := range fileList {
		method.File(model.Path).Delete()
	}

	return
}
