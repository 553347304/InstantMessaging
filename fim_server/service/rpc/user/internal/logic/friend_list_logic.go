package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src/sqls"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *user_rpc.FriendListRequest) (*user_rpc.FriendListResponse, error) {
	// todo: add your logic here and delete this line

	// l.svcCtx.DB.Preload("SendUserModel").Preload("ReceiveUserModel").Model(user_models.Friend{}).
	// 	Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId).Count(&total).
	// 	Find(&friends)

	friend := sqls.GetList(user_models.FriendModel{}, sqls.Mysql{
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", in.UserId, in.UserId),
		Preload: []string{"SendUserModel", "ReceiveUserModel"},
	})

	// 查哪些用户在线
	var list []*user_rpc.FriendInfo
	for _, fv := range friend.List {
		info := user_rpc.FriendInfo{}
		// 发起方
		if fv.SendUserId == uint(in.UserId) {
			info = user_rpc.FriendInfo{
				UserId: uint32(fv.ReceiveUserId),
				Name:   fv.ReceiveUserModel.Name,
				Avatar: fv.ReceiveUserModel.Avatar,
			}
		}
		// 接收方
		if fv.ReceiveUserId == uint(in.UserId) {
			info = user_rpc.FriendInfo{
				UserId: uint32(fv.SendUserId),
				Name:   fv.SendUserModel.Name,
				Avatar: fv.SendUserModel.Avatar,
			}
		}
		list = append(list, &info)
	}
	return &user_rpc.FriendListResponse{FriendList: list}, nil
}
