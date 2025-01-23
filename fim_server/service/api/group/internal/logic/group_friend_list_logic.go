package logic

import (
	"context"
	"fim_server/models/group_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupFriendListLogic {
	return &GroupFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupFriendListLogic) GroupFriendList(req *types.GroupFriendsListRequest) (resp *types.GroupFriendsListResponse, err error) {
	// todo: add your logic here and delete this line
	
	friendListResponse, err := l.svcCtx.UserRpc.Friend.FriendList(l.ctx, &user_rpc.ID{Id: uint32(req.UserId)})
	if err != nil {
		return nil, logs.Error(err)
	}
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ?", req.Id)
	var memberMap = map[uint]bool{}
	for _, model := range memberList {
		memberMap[model.UserId] = true
	}
	resp = new(types.GroupFriendsListResponse)
	for _, info := range friendListResponse.FriendList {
		resp.List = append(resp.List, types.GroupFriendsInfo{
			UserId:    uint(info.Id),
			Avatar:    info.Avatar,
			Name:      info.Name,
			IsInGroup: memberMap[uint(info.Id)],
		})
	}
	return
}
