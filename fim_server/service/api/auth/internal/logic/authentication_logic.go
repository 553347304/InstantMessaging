package logic

import (
	"context"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fim_server/utils/stores/valid"
	"fmt"
	
	"fim_server/service/api/auth/internal/svc"
	"fim_server/service/api/auth/internal/types"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp *types.AuthenticationResponse, err error) {
	// todo: add your logic here and delete this line
	
	if method.List(l.svcCtx.Config.WhiteList).InRegex(req.ValidPath) {
		logs.Info("白名单", req.ValidPath)
		return nil, nil
	}
	
	claims := valid.Jwt().Parse(req.Token)
	if claims == nil {
		return
	}
	
	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%s", req.Token)).Result()
	if err == nil {
		return nil, logs.Error("黑名单", err)
	}
	
	resp = &types.AuthenticationResponse{
		UserId: claims.PayLoad.UserId,
		Role:   claims.PayLoad.Role,
	}
	logs.Info("认证成功", resp)
	return resp, nil
}
