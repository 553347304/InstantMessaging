package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberDeleteLogic {
	return &GroupMemberDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberDeleteLogic) GroupMemberDelete(req *types.GroupMemberDeleteRequest) (resp *types.GroupMemberDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.Id, req.UserID).Error
	if err != nil || !(member.Role == 1 || member.Role == 2) {
		return nil, logs.Error("违规调用")
	}

	// 用户自己退群
	if req.UserID == req.MemberId {
		if member.Role == 1 {
			return nil, logs.Error("群主不能退群")
		}
		l.svcCtx.DB.Delete(&member)
		// 删除验证表记录
		l.svcCtx.DB.Create(&group_models.GroupValidModel{
			GroupId: member.GroupId,
			UserID:  req.UserID,
			Type:    2,
		})
		return
	}

	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.Id, req.MemberId).Error
	if err != nil {
		return nil, logs.Error("用户不是群成员", err.Error())
	}
	if member.Role >= member1.Role {
		return nil, logs.Error("不能踢自己/权限不足")
	}
	err = l.svcCtx.DB.Delete(&member1).Error
	if err != nil {
		return nil, logs.Error("移除失败", err.Error())
	}
	return
}
