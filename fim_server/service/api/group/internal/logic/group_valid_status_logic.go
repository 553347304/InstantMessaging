package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidStatusLogic {
	return &GroupValidStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidStatusLogic) GroupValidStatus(req *types.GroupValidStatusRequest) (resp *types.GroupValidStatusResponse, err error) {
	// todo: add your logic here and delete this line

	var groupValidModel group_models.GroupValidModel
	err = l.svcCtx.DB.Take(&groupValidModel, req.VaildId).Error
	if err != nil {
		return nil, logs.Error("验证记录不存在")
	}
	if groupValidModel.Status != 0 {
		return nil, logs.Error("已经处理过验证请求")
	}

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.UserId, groupValidModel.GroupId).Error
	if err != nil || !(member.Role == 1 || member.Role == 2) {
		return nil, logs.Error("权限不足")
	}

	switch req.Status {
	case 1: // 同意
	case 2:
	case 3:
	case 4:
	}

	l.svcCtx.DB.Model(&groupValidModel).UpdateColumn("status", req.Status)

	var isMember group_models.GroupMemberModel
	is := l.svcCtx.DB.Take(&isMember, "user_id = ? and group_id = ?", groupValidModel.UserId, groupValidModel.GroupId).Error
	if is == nil {
		return nil, logs.Error("该用户已经在群了")
	}

	// 将用户加到群里去
	l.svcCtx.DB.Create(&group_models.GroupMemberModel{
		GroupId: groupValidModel.GroupId,
		UserId:  groupValidModel.UserId,
		Role:    3,
	})

	return
}
