package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/encryption_and_decryptio/bcrypts"
	"fim_server/utils/encryption_and_decryptio/jwts"
	"fim_server/utils/stores/logs"
	
	"fim_server/service/api/auth/internal/svc"
	"fim_server/service/api/auth/internal/types"
	
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

	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "name = ?", req.Name).Error
	if err != nil {
		return nil, logs.Error("用户名错误")
	}

	if !bcrypts.Check(user.Password, req.Password) {
		return nil, logs.Error("密码错误")
	}

	token, err := jwts.GenToken(jwts.PayLoad{
		UserId: user.ID,
		Name:   user.Name,
		Role:   user.Role,
	})
	if err != nil {
		// l.svcCtx.ActionLogs.SetItem("error", err.Error())
		// l.svcCtx.ActionLogs.Err("服务内部错误")
		// l.svcCtx.RuntimeLogs.SetItemErr("xxx", err.Error())
		// err = errors.New("服务内部错误")
		return
	}
	
	// ctx := context.WithValue(l.ctx, "userID", fmt.Sprintf("%d", user.ID))
	// l.svcCtx.ActionLogs.Info("用户登录成功")
	// l.svcCtx.ActionLogs.SetCtx(ctx)

	logs.Info(l.ctx)
	
	return &types.LoginResponse{Token: token}, nil
}
