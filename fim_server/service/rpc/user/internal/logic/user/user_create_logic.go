package userlogic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

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
	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, "open_id = ?", in.OpenId).Error
	if err == nil {
		return nil, logs.Error("用户已经存在")
	}
	user = user_models.UserModel{
		Name:           in.Name,
		Avatar:         in.Avatar,
		Role:           in.Role,
		OpenId:         in.OpenId,
		RegisterSource: in.RegisterSource,
	}

	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logx.Error(err)
		return nil, logs.Error("创建用户失败")
	}

	// 创建用户配置
	l.svcCtx.DB.Create(&user_models.UserConfigModel{
		UserID:        user.ID,
		RecallMessage: nil,
		FriendOnline:  false,
		Sound:         true,
		SecureLink:    false,
		SavePassword:  false,
		SearchUser:    2,
		Valid:         2,
		Online:        true,
	})

	return &user_rpc.UserCreateResponse{UserID: int32(user.ID)}, nil
}
