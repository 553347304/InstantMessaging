package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBaseInfoLogic {
	return &UserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserBaseInfoLogic) UserBaseInfo(in *user_rpc.UserBaseInfoRequest) (*user_rpc.UserBaseInfoResponse, error) {
	// todo: add your logic here and delete this line

	var user user_models.UserModel
	err := l.svcCtx.DB.Preload("UserConfigModel").Take(&user, in.UserId).Error
	if err != nil {
		return nil, logs.Error("用户不存在", in.UserId)
	}

	return &user_rpc.UserBaseInfoResponse{
		UserId: in.UserId,
		Name:   user.Name,
		Avatar: user.Avatar,
	}, nil
}
