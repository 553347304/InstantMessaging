package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/models/user_models"
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupTopLogic {
	return &GroupTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupTopLogic) GroupTop(req *types.GroupTopRequest) (resp *types.GroupTopResponse, err error) {
	// todo: add your logic here and delete this line
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.GroupId, req.UserId).Error
	if err != nil {
		return nil, logs.Error("用户不是群成员", err.Error())
	}

	var userModel user_models.UserModel
	l.svcCtx.DB.Take(&userModel, req.UserId)

	index := method.List(userModel.Top.GroupId).In(req.GroupId)

	// 置顶
	if req.Top && index == -1 {
		userModel.Top.GroupId = append(userModel.Top.GroupId, req.GroupId)
		userModel.Top.GroupId = method.List(userModel.Top.GroupId).Sort(true)
		l.svcCtx.DB.Updates(&userModel)
	}

	// 取消置顶
	if !req.Top && index != -1 {
		userModel.Top.GroupId = method.List(userModel.Top.GroupId).Delete(index)
		userModel.Top.GroupId = method.List(userModel.Top.GroupId).Sort(true)
		l.svcCtx.DB.Updates(&userModel)
	}

	return
}
