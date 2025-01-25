package userlogic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/conv"

	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user_rpc.IdList) (*user_rpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	var userList = make([]user_models.UserModel, 0)
	l.svcCtx.DB.Preload("UserConfigModel").Find(&userList, in.Id)

	resp := new(user_rpc.UserInfoResponse)

	resp.InfoList = make(map[uint32]*user_rpc.UserInfo)
	for i, user := range userList {
		info := conv.Struct(user_rpc.UserInfo{}).Type(user)
		if i == 0 {
			resp.Info = &info
		}
		resp.InfoList[uint32(user.ID)] = &info
		resp.Total = int64(i)
	}
	if resp.Total == 0 {
		return resp, conv.Type("用户服务错误").Error()
	}

	return resp, nil
}
