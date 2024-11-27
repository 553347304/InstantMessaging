package logic

import (
	"context"
	"errors"
	"fim_server/fim_user/user_models"

	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	// todo: add your logic here and delete this line

	var user user_models.User
	err := l.svcCtx.DB.Take(&user, "open_id = ?", in.OpenId).Error
	if err == nil {
		return nil, errors.New("用户已经存在")
	}
	user = user_models.User{
		Name:           in.Name,
		Avatar:         in.Avatar,
		Role:           int8(in.Role),
		OpenId:         in.OpenId,
		RegisterSource: in.RegisterSource,
	}

	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建用户失败")
	}

	// 创建用户配置
	l.svcCtx.DB.Create(&user_models.UserConfig{
		UserId:        user.ID,
		RecallMessage: nil,
		FriendOnline:  false,
		Sound:         true,
		SecureLink:    false,
		SavePassword:  false,
		SearchUser:    2,
		Auth:          2,
		Online:        true,
	})

	return &user_rpc.UserCreateResponse{UserId: int32(user.ID)}, nil
}
