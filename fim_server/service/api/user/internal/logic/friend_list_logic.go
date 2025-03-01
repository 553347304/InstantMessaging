package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
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

	// 获取好友列表
	friend := src.Mysql(src.ServiceMysql[user_models.FriendModel]{
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId),
		Preload: []string{"SendUserModel", "ReceiveUserModel"},
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}).GetList()

	// 查在线用户
	onlineMap := l.svcCtx.Redis.HGetAll("user_online").Val()
	var onlineUserMap = map[uint64]bool{}
	for key, _ := range onlineMap {
		v, err := conv.Type(key).Uint64()
		if err != nil {
			continue
		}
		onlineUserMap[v] = true
	}

	var list []types.FriendInfoResponse
	for _, fv := range friend.List {
		// 发起方
		info := types.FriendInfoResponse{}
		if fv.SendUserId == req.UserId {
			info = types.FriendInfoResponse{
				UserId:   fv.ReceiveUserId,
				Username:     fv.ReceiveUserModel.Username,
				Sign:     fv.ReceiveUserModel.Sign,
				Avatar:   fv.ReceiveUserModel.Avatar,
				Notice:   fv.SendUserNotice,
				IsOnline: onlineUserMap[fv.ReceiveUserId],
			}
		}
		// 接收方
		if fv.ReceiveUserId == req.UserId {
			info = types.FriendInfoResponse{
				UserId:   fv.SendUserId,
				Username:     fv.SendUserModel.Username,
				Sign:     fv.SendUserModel.Sign,
				Avatar:   fv.SendUserModel.Avatar,
				Notice:   fv.ReceiveUserNotice,
				IsOnline: onlineUserMap[fv.SendUserId],
			}
		}
		list = append(list, info)
	}
	return &types.FriendListResponse{List: list, Total: friend.Total}, nil
}
