package friendlogic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/src"

	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

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

func (l *FriendListLogic) FriendList(in *user_rpc.ID) (*user_rpc.FriendListResponse, error) {
	// todo: add your logic here and delete this line

	friend := src.Mysql(src.ServiceMysql[user_models.FriendModel]{
		Model:   user_models.FriendModel{},
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", in.Id, in.Id),
		Preload: []string{"SendUserModel", "ReceiveUserModel"},
	}).GetList()

	// 查哪些用户在线
	var list []*user_rpc.UserInfo
	for _, fv := range friend.List {
		info := user_rpc.UserInfo{}
		// 发起方
		if fv.SendUserId == in.Id {
			info = user_rpc.UserInfo{
				Id:     fv.ReceiveUserId,
				Username:   fv.ReceiveUserModel.Username,
				Avatar: fv.ReceiveUserModel.Avatar,
			}
		}
		// 接收方
		if fv.ReceiveUserId == in.Id {
			info = user_rpc.UserInfo{
				Id:     fv.SendUserId,
				Username:   fv.SendUserModel.Username,
				Avatar: fv.SendUserModel.Avatar,
			}
		}
		list = append(list, &info)
	}

	return &user_rpc.FriendListResponse{FriendList: list}, nil
}
