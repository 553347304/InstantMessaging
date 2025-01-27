package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDeleteLogic {
	return &GroupDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupDeleteLogic) GroupDelete(req *types.GroupDeleteRequest) (resp *types.GroupDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	var groupMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Preload("GroupModel").Take(&groupMember, "group_id = ? and user_id = ?", req.Id, req.UserID).Error
	if err != nil {
		return nil, logs.Error("群不存在或用户不是群成员", err.Error())
	}
	if groupMember.Role != 1 {
		return nil, logs.Error("只有群主才能解散")
	}

	// 关联删除
	var messageList []group_models.GroupMessageModel
	l.svcCtx.DB.Find(&messageList, "group_id = ?", req.Id).Delete(&messageList)
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ?", req.Id).Delete(&memberList)
	var vList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&vList, "group_id = ?", req.Id).Delete(&vList)
	var group group_models.GroupModel
	l.svcCtx.DB.Find(&group, req.Id).Delete(&group)
	logs.Info("删除群", group.Name)
	logs.Info("群成员数", len(memberList))
	logs.Info("群消息数", len(messageList))
	logs.Info("群验证消息", len(vList))
	return
}
