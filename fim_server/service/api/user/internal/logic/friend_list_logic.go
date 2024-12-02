package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/src/sqls"
	"fim_server/utils/stores"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

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
	// l.svcCtx.DB.Preload("SendUserModel").Preload("ReceiveUserModel").Model(user_models.Friend{}).
	// 	Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId).Count(&total).
	// 	Find(&friends)

	friends, total := sqls.GetList(user_models.FriendModel{}, sqls.Mysql{
		DB:      l.svcCtx.DB,
		Preload: []string{"SendUserModel", "ReceiveUserModel"},
		PageInfo: stores.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	var list []types.FriendInfoResponse
	for _, fv := range friends {
		// 发起方
		info := types.FriendInfoResponse{}
		if fv.SendUserId == req.UserId {
			info = types.FriendInfoResponse{
				UserId: fv.SendUserId,
				Name:   fv.SendUserModel.Name,
				Sign:   fv.SendUserModel.Sign,
				Avatar: fv.SendUserModel.Avatar,
				Notice: fv.SendUserNotice,
			}
		}
		// 接收方
		if fv.SendUserId == req.UserId {
			info = types.FriendInfoResponse{
				UserId: fv.ReceiveUserId,
				Name:   fv.ReceiveUserModel.Name,
				Sign:   fv.ReceiveUserModel.Sign,
				Avatar: fv.ReceiveUserModel.Avatar,
				Notice: fv.ReceiveUserNotice,
			}
		}
		list = append(list, info)
	}
	return &types.FriendListResponse{List: list, Total: int(total)}, nil
}
