package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberNameLogic {
	return &GroupMemberNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberNameLogic) GroupMemberName(req *types.GroupMemberNameRequest) (resp *types.GroupMemberNameResponse, err error) {
	// todo: add your logic here and delete this line

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.Id, req.UserId).Error
	if err != nil || !(member.Role == 1 || member.Role == 2) {
		return nil, logs.Error("违规调用")
	}
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.Id, req.MemberId).Error
	if err != nil {
		return nil, logs.Error("用户不是群成员", err.Error())
	}

	if member.Role >= member1.Role && req.UserId != req.MemberId {
		return nil, logs.Error("权限不足")
	}
	logs.Info(member1)
	l.svcCtx.DB.Model(&member1).Updates(map[string]any{"role": req.Name})
	return
}
