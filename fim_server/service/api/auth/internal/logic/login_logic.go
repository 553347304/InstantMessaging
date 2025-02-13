package logic

import (
	"context"
	"fim_server/models/mtype"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/valid"
	
	"fim_server/service/api/auth/internal/svc"
	"fim_server/service/api/auth/internal/types"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "name = ?", req.Name).Error
	if err != nil {
		l.svcCtx.RpcLog.Info(l.ctx, "用户名错误: "+req.Name)
		return nil, logs.Error("用户名错误")
	}

	
	if !valid.Bcrypt().Check(user.Password, req.Password) {
		l.svcCtx.RpcLog.Info(l.ctx, "密码错误: "+req.Password)
		return nil, logs.Error("密码错误")
	}

	token := valid.Jwt().Hash(mtype.PayLoad{
		UserID: user.ID,
		Name:   user.Name,
		Role:   user.Role,
	})
	if token == "" {
		l.svcCtx.RpcLog.Info(l.ctx, "用户登录成功")
		return nil, logs.Error("登录服务内部错误")
	}
	l.svcCtx.RpcLog.Info(l.ctx, "用户登录成功")

	return &types.LoginResponse{Token: token}, nil
}
