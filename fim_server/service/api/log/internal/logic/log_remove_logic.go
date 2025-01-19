package logic

import (
	"context"
	
	"fim_server/service/api/log/internal/svc"
	"fim_server/service/api/log/internal/types"
)

type LogRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogRemoveLogic {
	return &LogRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogRemoveLogic) LogRemove(req *types.LogRemoveRequest) (resp *types.Empty, err error) {
	// var logList []logs_model.LogModel
	// l.svcCtx.DB.Find(&logList, "id in ?", req.IdList)
	// if len(logList) > 0 {
	// 	l.svcCtx.DB.Delete(&logList)
	// 	l.svcCtx.ActionLogs.SetItem("删除日志条数", len(logList))
	// 	logx.Infof("删除日志条数 %d", len(logList))
	// }
	// l.svcCtx.ActionLogs.Info("删除日志操作")
	// l.svcCtx.ActionLogs.Save(l.ctx)
	// 
	return
}
