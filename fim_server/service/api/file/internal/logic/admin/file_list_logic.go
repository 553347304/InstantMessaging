package admin

import (
	"context"
	"fim_server/models/file_models"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	
	"fim_server/service/api/file/internal/svc"
	"fim_server/service/api/file/internal/types"
)

type FileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.PageInfo) (resp *types.FileListResponse, err error) {
	// todo: add your logic here and delete this line
	
	fileResponse := src.Mysql(src.ServiceMysql[file_models.FileModel]{
		DB:       l.svcCtx.DB,
		PageInfo: conv.Struct(src.PageInfo{}).Type(req),
		Where:    "name = ?",
	}).GetList()
	resp = new(types.FileListResponse)
	resp = conv.Struct(&types.FileListResponse{}).Type(fileResponse)
	return
}
