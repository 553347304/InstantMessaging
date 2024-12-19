package logic

import (
	"context"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupFriendsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupFriendsListLogic {
	return &GroupFriendsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupFriendsListLogic) GroupFriendsList(req *types.GroupFriendsListRequest) (resp *types.GroupFriendsListResponse, err error) {
	// todo: add your logic here and delete this line


	return
}
