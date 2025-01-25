package Admin

import (
	"context"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.RequestDelete) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
