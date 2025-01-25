package logic

import (
	"context"
	"fim_server/models/log_models"
	"fim_server/service/server/response"
	"fim_server/utils/src"

	"fim_server/service/api/log/internal/svc"
	"fim_server/service/api/log/internal/types"
)

type LogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogListLogic {
	return &LogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogListLogic) LogList(req *types.PageInfo) (resp *response.List[log_models.LogModel], err error) {

	log := src.Mysql(src.ServiceMysql[log_models.LogModel]{
		DB:    l.svcCtx.DB,
		Where: "name like ? or type like ?",
		PageInfo: src.PageInfo{
			Key:   req.Key,
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
	}).GetList()
	// Likes: []string{"ip", "user_nickname", "title"},
	resp = new(response.List[log_models.LogModel])
	resp.List = log.List
	resp.Total = log.Total
	return
}
