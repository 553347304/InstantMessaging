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
	if req.Status != 1 {
		return nil, logs.Error("已经处理过验证请求")
	}

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.UserId, groupValidModel.GroupId).Error
	if err != nil || !(member.Role == 1 || member.Role == 2) {
		return nil, logs.Error("权限不足")
	}

	switch req.Status {
	case 1: // 同意
		return
	case 2:
		friendVerify.ReceiveStatus = 2 // 拒绝
	case 3:
		friendVerify.ReceiveStatus = 3 // 忽略
	case 4:
		friendVerify.Status = 4 // 删除
		l.svcCtx.DB.Delete(&friendVerify)
		return nil, nil
	}
	friendVerify.Status = friendVerify.ReceiveStatus
	logs.Info(friendVerify.Status)
	l.svcCtx.DB.Model(&friendVerify).Updates(map[string]any{
		"status":         friendVerify.Status,
		"receive_status": friendVerify.ReceiveStatus,
	})

	return
}
