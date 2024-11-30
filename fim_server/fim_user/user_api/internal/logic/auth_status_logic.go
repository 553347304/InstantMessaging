package logic

import (
	"context"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthStatusLogic {
	return &AuthStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthStatusLogic) AuthStatus(req *types.FriendAuthStatusRequest) (resp *types.FriendAuthStatusResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
