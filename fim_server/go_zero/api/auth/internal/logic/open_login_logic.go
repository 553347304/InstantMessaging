package logic

import (
	"context"
	"fim_server/go_zero/rpc/user/user_rpc"
	"fim_server/models/user_models"
	"fim_server/utils/encryption_and_decryptio/jwts"
	"fim_server/utils/open_login"
	"fim_server/utils/stores/logs"
	"fmt"

	"fim_server/go_zero/api/auth/internal/svc"
	"fim_server/go_zero/api/auth/internal/types"

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

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	type OpenInfo struct {
		Name   string
		OpenId string
		Avatar string
	}

	var info OpenInfo
	switch req.Flag {
	case "qq":
		qqInfo, errs := open_login.NewQQLogin(req.Code)

		info = OpenInfo{
			OpenId: qqInfo.OpenID,
			Name:   qqInfo.Name,
			Avatar: qqInfo.Avatar,
		}
		err = errs
	default:

		err = logs.Error("不支持的第三方登录")
	}

	if err != nil {
		logs.Error("登录失败", err)
		return nil, logs.Error("登录失败")
	}
	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenId).Error
	if err != nil {
		fmt.Println("注册服务")
		result, errs := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
			Name:           info.Name,
			Password:       "",
			Role:           2,
			Avatar:         info.Avatar,
			OpenId:         info.OpenId,
			RegisterSource: "qq",
		})
		if errs != nil {
			return nil, logs.Error("登录失败")
		}
		user.Model.ID = uint(result.UserId)
		user.Role = 2
		user.Name = info.Name
	}
	// 登录
	token, err := jwts.GenToken(jwts.PayLoad{
		UserId: user.ID,
		Name:   user.Name,
		Role:   user.Role,
	})
	if err != nil {
		logx.Error(err)
		return nil, logs.Error("服务器内部错误")
	}
	return &types.LoginResponse{Token: token}, nil

}
