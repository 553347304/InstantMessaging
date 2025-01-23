package Admin

import (
	"context"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
	
)

type GroupMessageDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMessageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMessageDeleteLogic {
	return &GroupMessageDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMessageDeleteLogic) GroupMessageDelete(req *types.RequestDelete) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
