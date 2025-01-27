package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberAddLogic {
	return &GroupMemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberAddLogic) GroupMemberAdd(req *types.GroupMemberAddRequest) (resp *types.GroupMemberAddResponse, err error) {
	// todo: add your logic here and delete this line

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Preload("GroupModel").Take(&member, "group_id = ? and user_id = ?", req.Id, req.UserID).Error
	if err != nil {
		return nil, logs.Error("违规调用")
	}
	if member.Role == 3 {
		if !member.GroupModel.IsInvite {
			return nil, logs.Error("管理员未开放好友邀请入群功能")
		}
	}

	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ? and user_id in ?", req.Id, req.MemberIdList)
	if len(memberList) > 0 {
		return nil, logs.Error("已经有用户在群里了")
	}
	for _, memberId := range req.MemberIdList {
		memberList = append(memberList, group_models.GroupMemberModel{
			GroupId: req.Id,
			UserID:  memberId,
			Role:    3,
		})
	}
	err = l.svcCtx.DB.Create(&memberList).Error
	if err != nil {
		return nil, logs.Error("邀请失败")
	}
	return
}
