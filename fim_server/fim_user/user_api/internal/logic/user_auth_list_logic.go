package logic

import (
	"context"
	"fim_server/fim_user/user_models"
	"fim_server/utils/stores"
	"fim_server/utils/stores/server/sqls"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"

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

	fvs, total := sqls.GetList(user_models.FriendAuthModel{}, sqls.Mysql{
		DB:      l.svcCtx.DB,
		Preload: []string{"ReceiveUserModel.UserConfigModel"},
		PageInfo: stores.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})

	var list []types.FriendAuthInfoResponse
	for _, fv := range fvs {
		list = append(list, types.FriendAuthInfoResponse{
			UserId:       fv.ReceiveUserId,
			Name:         fv.ReceiveUserModel.Name,
			Avatar:       fv.ReceiveUserModel.Avatar,
			AuthMessage:  fv.AuthMessage,
			AuthQuestion: (*types.AuthQuestion)(fv.AuthQuestion),
			Status:       fv.Status,
			Auth:         fv.ReceiveUserModel.UserConfigModel.Auth,
			Id:           fv.ID,
		})
	}

	return &types.FriendAuthResponse{List: list, Total: total}, nil
}
