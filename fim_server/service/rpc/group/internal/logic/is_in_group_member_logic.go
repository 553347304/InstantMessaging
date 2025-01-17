package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/group/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsInGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsInGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsInGroupMemberLogic {
	return &IsInGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsInGroupMemberLogic) IsInGroupMember(in *group_rpc.IsInGroupMemberRequest) (resp *group_rpc.EmptyResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(group_rpc.EmptyResponse)
	// 判断用户是否在群里
	var groupModel group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&groupModel, "group_id = ? and user_id = ?", in.GroupId, in.UserId).Error
	if err != nil {
		return nil, logs.Error("不在群里面")
	}

	return &group_rpc.EmptyResponse{}, nil
}
