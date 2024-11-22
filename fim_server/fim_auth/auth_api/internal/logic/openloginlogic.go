package logic

import (
	"context"
	"errors"
	"fim_server/fim_auth/auth_api/internal/svc"
	"fim_server/fim_auth/auth_api/internal/types"
	"fim_server/fim_auth/auth_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils/open_login"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginInfoRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	switch req.Flag {
	case "qq":
		info, err := open_login.NewQQLogin(req.Code)
		if err != nil {
			logx.Error(err)
			return nil, errors.New("登录失败")
		}
		var user auth_models.User
		err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
		if err != nil {
			fmt.Println("注册服务")
			result, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
				Name:     info.Name,
				Password: "",
				Role:     2,
				Avatar:   info.Avatar,
				OpenId:   info.OpenID,
			})
			if err != nil {
				return nil, errors.New("登录失败")
			}
			user.Model.ID = uint(result.UserId)
			user.Role = 2
			user.Name = info.Name
		}
		// 登录
		jwts.
		// jwts.GenToken()
	}

	return
}
