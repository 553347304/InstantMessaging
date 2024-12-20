package logic

import (
	"context"
	"fim_server/models"
	"fim_server/models/group_models"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUpdateLogic {
	return &GroupUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUpdateLogic) GroupUpdate(req *types.GroupUpdateRequest) (resp *types.GroupUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.GroupUpdateResponse)
	var groupMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Preload("GroupModel").Take(&groupMember, "group_id = ? and user_id = ?", req.Id, req.UserId).Error
	if err != nil {
		return nil, logs.Error("群不存在或用户不是群成员", err.Error())
	}
	if !(groupMember.Role == 1 || groupMember.Role == 2) {
		return nil, logs.Error("只能是群主或管理员才能更新")
	}
	groupMaps := conv.StructMap(*req, "conf")
	if len(groupMaps) != 0 {
		_, ok := groupMaps["auth_question"]
		if ok {
			delete(groupMaps, "auth_question")
			l.svcCtx.DB.Model(&groupMember.GroupModel).Updates(&group_models.GroupModel{
				VerifyInfo: models.VerifyInfo{
					Issue:  req.VerifyInfo.Issue,
					Answer: req.VerifyInfo.Answer,
				},
			})
		}
		err = l.svcCtx.DB.Model(&groupMember.GroupModel).Updates(groupMaps).Error
		if err != nil {
			return nil, logs.Error("用户信息更新失败")
		}
	}
	return
}
