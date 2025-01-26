package logic

import (
	"context"
	"fim_server/models/setting_models"
	"fim_server/models/user_models"
	"fim_server/service/rpc/setting/setting_rpc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/encryption_and_decryptio/jwts"
	"fim_server/utils/open_api/open_api_qq"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fmt"
	
	"fim_server/service/api/auth/internal/svc"
	"fim_server/service/api/auth/internal/types"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type OpenInfo struct {
	OpenID string
	Name   string
	Avatar string
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	var conf setting_models.ConfigModel
	settingRpc, _ := l.svcCtx.SettingRpc.SettingInfo(l.ctx, &setting_rpc.Empty{})
	if !conv.Json().Unmarshal(settingRpc.Data, &conf) {
		return nil, logs.Error("系统配置服务异常")
	}
	
	var info OpenInfo
	switch req.Flag {
	case "qq":
		r := open_api_qq.Login(open_api_qq.LoginConfig{
			Code:     req.Code,
			AppID:    conf.OpenLogin.QQ.AppID,
			AppKey:   conf.OpenLogin.QQ.Key,
			Redirect: conf.OpenLogin.QQ.Redirect,
		})
		if r.Error != nil {
			return nil, r.Error
		}
		logs.Info(r)
		return
		// info = OpenInfo{
		// 	OpenID: qqInfo.OpenID,
		// 	Name:   qqInfo.Name,
		// 	Avatar: qqInfo.Avatar,
		// }
		// err = openError
	default:
		err = logs.Error("不支持的第三方登录")
	}
	
	if err != nil {
		logs.Error("登录失败", err)
		return nil, logs.Error("登录失败")
	}
	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		fmt.Println("注册服务")
		result, errs := l.svcCtx.UserRpc.User.UserCreate(l.ctx, &user_rpc.UserCreateRequest{
			Name:           info.Name,
			Password:       "",
			Role:           2,
			Avatar:         info.Avatar,
			OpenId:         info.OpenID,
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
