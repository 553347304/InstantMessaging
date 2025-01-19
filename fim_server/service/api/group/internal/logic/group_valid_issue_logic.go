package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidIssueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidIssueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidIssueLogic {
	return &GroupValidIssueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidIssueLogic) GroupValidIssue(req *types.GroupValidIssueRequest) (resp *types.GroupValidIssueResponse, err error) {
	// todo: add your logic here and delete this line
	
	var groupModel group_models.GroupModel
	err = l.svcCtx.DB.Take(&groupModel, "id = ?", req.Id).Error
	if err != nil {
		return nil, logs.Error("群不存在")
	}
	resp = new(types.GroupValidIssueResponse)
	resp.Valid = groupModel.Valid
	resp.ValidInfo = conv.Struct(types.ValidInfo{}).Type(groupModel.ValidInfo)
	return
}
