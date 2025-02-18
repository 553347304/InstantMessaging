package logic

import (
	"context"
	"fim_server/models/setting_models"
	"fim_server/models/user_models"
	"fim_server/service/rpc/setting/setting_rpc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/open_api/open_api_qq"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/valid"
	
	"fim_server/service/api/auth/internal/svc"
	"fim_server/service/api/auth/internal/types"
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
	OpenId   string
	Username string
	Avatar   string
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
		info.Username = r.Nickname
		info.OpenId = r.OpenID
		info.Avatar = r.Avatar
	default:
		return nil, logs.Error("不支持的第三方登录")
	}
	
	
	// 注册用户
	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenId).Error
	if err != nil {
		result, errs := l.svcCtx.UserRpc.User.UserCreate(l.ctx, &user_rpc.UserCreateRequest{
			Username:       info.Username,
			Password:       "",
			Role:           2,
			Avatar:         info.Avatar,
			OpenId:         info.OpenId,
			RegisterSource: "qq",
		})
		if errs != nil {
			return nil, logs.Error("注册失败")
		}
		user.ID = result.UserId
		user.Role = 2
		user.Username = info.Username
	}
	
	// 登录
	token := valid.Jwt().Hash(valid.PayLoad{
		UserId:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	if token == "" {
		return nil, logs.Error("服务器内部错误")
	}
	return &types.LoginResponse{Token: token}, nil
	
}
