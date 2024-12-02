package logic

import (
	"context"
	"fim_server/utils/encryption_and_decryptio/jwts"
	"fmt"
	"time"

	"fim_server/service/api/auth/internal/svc"
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

	key := fmt.Sprintf("logout_%s", token)
	l.svcCtx.Redis.SetNX(key, "", expiration)

	resp = "注销成功"
	return
}
