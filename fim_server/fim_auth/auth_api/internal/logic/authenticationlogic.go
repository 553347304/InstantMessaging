package logic

import (
	"context"
	"errors"
	"fim_server/fim_auth/auth_api/internal/svc"
	"fim_server/utils/jwts"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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

func (l *AuthenticationLogic) Authentication(r *http.Request) (resp string, err error) {
	// todo: add your logic here and delete this line

	claims := jwts.ParseToken(r.Header.Get("token"))
	if claims == nil {
		return
	}
	result, err := l.svcCtx.Redis.Get(fmt.Sprintf("logout_%d", claims.UserId)).Result()
	if err != nil {
		return "", errors.New("认证失败")
	}
	logx.Error(result)

	resp = "ok"
	err = nil
	return
}
