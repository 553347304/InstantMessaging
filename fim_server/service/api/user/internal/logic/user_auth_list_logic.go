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

type UserAuthListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAuthListLogic {
	return &UserAuthListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAuthListLogic) UserAuthList(req *types.FriendAuthRequest) (resp *types.FriendAuthResponse, err error) {
	// todo: add your logic here and delete this line

	fvs, total := sqls.GetList(user_models.FriendAuthModel{ReceiveUserId: req.UserId}, sqls.Mysql{
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId),
		Preload: []string{"ReceiveUserModel.UserConfigModel", "SendUserModel.UserConfigModel"},
		PageInfo: stores.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})

	var list []types.FriendAuthInfo
	for _, fv := range fvs {
		info := types.FriendAuthInfo{
			AuthMessage:  fv.AuthMessage,
			AuthQuestion: (*types.AuthQuestion)(fv.AuthQuestion),
			Status:       fv.Status,
			Id:           fv.ID,
		}

		if fv.SendUserId == req.UserId {
			// 发起方
			info.UserId = fv.SendUserId
			info.Name = fv.SendUserModel.Name
			info.Avatar = fv.SendUserModel.Avatar
			info.Auth = fv.SendUserModel.UserConfigModel.Auth
			info.Status = fv.SendStatus
			info.Flag = "send"
		}
		if fv.ReceiveUserId == req.UserId {
			// 接收方
			info.UserId = fv.ReceiveUserId
			info.Name = fv.ReceiveUserModel.Name
			info.Avatar = fv.ReceiveUserModel.Avatar
			info.Auth = fv.ReceiveUserModel.UserConfigModel.Auth
			info.Status = fv.ReceiveStatus
			info.Flag = "receive"
		}

		list = append(list, info)
	}

	return &types.FriendAuthResponse{List: list, Total: total}, nil
}
