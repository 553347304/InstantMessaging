package logic

import (
	"context"
	"fim_server/fim_user/user_models"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {
	// todo: add your logic here and delete this line

	var friends []user_models.Friend
	l.svcCtx.DB.Preload("SendUserId").Preload("ReceiveUserModel").
		Find(&friends, "send_user_id = ? or receive_user_id = ?", req.UserId)
	var list []types.FriendInfoResponse
	for _, friend := range friends {
		// 发起方
		info := types.FriendInfoResponse{}
		if friend.SendUserId == req.UserId {
			info = types.FriendInfoResponse{
				UserId: friend.SendUserId,
				Name:   friend.SendUserModel.Name,
				Sign:   friend.SendUserModel.Sign,
				Avatar: friend.SendUserModel.Avatar,
				Notice: friend.SendUserNotice,
			}
		}
		// 接收方
		if friend.SendUserId == req.UserId {
			info = types.FriendInfoResponse{
				UserId: friend.ReceiveUserId,
				Name:   friend.ReceiveUserModel.Name,
				Sign:   friend.ReceiveUserModel.Sign,
				Avatar: friend.ReceiveUserModel.Avatar,
				Notice: friend.ReceiveUserNotice,
			}
		}
		list = append(list, info)
	}
	return
}
