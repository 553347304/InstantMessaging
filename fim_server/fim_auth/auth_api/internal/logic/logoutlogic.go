package logic

import (
	"context"
	"fim_server/utils/jwts"
	"fmt"
	"time"

	"fim_server/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp string, err error) {
	// todo: add your logic here and delete this line

	claims := jwts.ParseToken(token)
	if claims == nil {
		return
	}
	now := time.Now()
	expiration := claims.ExpiresAt.Time.Sub(now)

	l.svcCtx.Redis.SetNX(fmt.Sprintf("logout_%d", claims.UserId), "", expiration)
	resp = "注销成功"
	return
}
