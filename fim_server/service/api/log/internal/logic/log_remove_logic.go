package logic

import (
	"context"
	log_model "fim_server/models/log_models"
	"fim_server/utils/stores/logs"

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
	var logList []log_model.LogModel
	l.svcCtx.DB.Find(&logList, "id in ?", req.IdList)
	if len(logList) > 0 {
		l.svcCtx.DB.Delete(&logList)
	}
	l.svcCtx.RpcLog.Info(l.ctx, logs.Info("删除日志条数", len(logList)))
	return
}
