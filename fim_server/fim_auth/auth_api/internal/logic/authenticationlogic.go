package logic

import (
	"context"
	"errors"
	"fim_server/fim_auth/auth_api/internal/svc"
	"fim_server/fim_auth/auth_api/internal/types"
	"fim_server/utils/jwts"
	"fim_server/utils/stores/algorithms"
	"fim_server/utils/stores/logs"
	"fmt"

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

	if algorithms.InList(l.svcCtx.Config.WhiteList, req.ValidPath) {
		logs.Info("白名单", req.ValidPath)
		return nil, nil
	}

	claims := jwts.ParseToken(req.Token)
	if claims == nil {
		err = errors.New("认证失败: " + req.Token)
		return
	}

	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%s", req.Token)).Result()
	if err == nil {
		return nil, logs.Error("黑名单", err.Error())
	}

	resp = &types.AuthenticationResponse{
		UserId: claims.UserId,
		Role:   claims.Role,
	}
	logs.Info("认证成功", resp)
	return resp, nil
}
