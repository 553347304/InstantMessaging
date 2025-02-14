package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *user_models.UserModel, err error) {
	// todo: add your logic here and delete this line
	
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: []uint64{req.UserId}})
	if err != nil {
		return nil, logs.Error(err.Error())
	}
	var user user_models.UserModel
	if !method.Struct().To(userResponse.Info, &user) {
		return nil, logs.Error("转换失败")
	}
	resp = &user
	
	method.Struct().Delete(resp, "Password")
	
	return resp, nil
}
