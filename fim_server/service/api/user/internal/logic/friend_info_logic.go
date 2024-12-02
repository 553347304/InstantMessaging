package logic

import (
	"context"
	"encoding/json"
	"fim_server/models/user_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	// todo: add your logic here and delete this line

	var friend user_models.FriendModel
	if !friend.IsFriend(l.svcCtx.DB, req.UserId, req.FriendId) {
		return nil, logs.Error("他不是你的好友")
	}

	result, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.FriendId),
	})

	if err != nil {
		return nil, logs.Error(err)
	}
	var friendUser user_models.UserModel
	err = json.Unmarshal(result.Data, &friendUser)
	if err != nil {
		return nil, logs.Error("绑定失败", err)
	}

	response := &types.FriendInfoResponse{
		UserId: friendUser.ID,
		Name:   friendUser.Name,
		Avatar: friendUser.Avatar,
		Notice: friend.GetUserNotice(req.UserId),
	}

	return response, nil
}
