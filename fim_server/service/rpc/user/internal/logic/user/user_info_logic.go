package userlogic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/method"
	
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(in *user_rpc.IdList) (*user_rpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var userList = make([]user_models.UserModel, 0)
	l.svcCtx.DB.Preload("UserConfigModel").Find(&userList, in.Id)
	
	resp := new(user_rpc.UserInfoResponse)
	
	resp.InfoList = make(map[uint64]*user_rpc.UserInfo)
	for i, u := range userList {
		var user user_rpc.UserInfo
		method.Struct().To(u, &user)
		if i == 0 {
			resp.Info = &user
		}
		resp.InfoList[u.ID] = &user
	}
	resp.Total = int64(len(resp.InfoList))
	if resp.Total == 0 {
		return resp, conv.Type("用户服务错误").Error()
	}
	
	return resp, nil
}
