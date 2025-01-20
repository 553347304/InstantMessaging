package logic

import (
	"context"
	"fim_server/models/log_models"
	"fim_server/service/server/response"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"
	
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
	log := sqls.GetList(log_models.LogModel{}, sqls.Mysql{
		DB: l.svcCtx.DB,
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
	})
	// Likes: []string{"ip", "user_nickname", "title"},
	resp = new(response.List[log_models.LogModel])
	
	resp.List = log.List
	resp.Total = log.Total
	return
}
