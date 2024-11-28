package logic

import (
	"context"
	"encoding/json"
	"fim_server/fim_user/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"

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

func (l *UserInfoLogic) UserInfo(in *user_rpc.UserInfoRequest) (*user_rpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	var user user_models.User
	err := l.svcCtx.DB.Preload("UserConfig").Take(&user, in.UserId).Error
	if err != nil {
		return nil, logs.Error("用户不存在", in.UserId)
	}
	byteData, _ := json.Marshal(user)

	return &user_rpc.UserInfoResponse{Data: byteData}, nil
}
