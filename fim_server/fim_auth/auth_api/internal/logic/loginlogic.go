package logic

import (
	"context"
	"errors"
	"fim_server/fim_auth/auth_api/internal/svc"
	"fim_server/fim_auth/auth_api/internal/types"
	"fim_server/fim_auth/auth_models"
	"fim_server/utils/bcrypts"
	"fim_server/utils/jwts"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	var user auth_models.User
	err = l.svcCtx.DB.Take(&user, req.Username).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}

	if bcrypts.Check(user.Password, req.Password) {
		err = errors.New("用户名或密码错误")
		return
	}

	token, err := jwts.GenToken(jwts.PayLoad{
		UserId: user.ID,
		Name:   user.Name,
		Role:   user.Role,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.LoginResponse{Token: token}, nil
}
