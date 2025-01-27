package grouplogic

import (
	"context"

	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/group/internal/svc"
)

type UserGroupSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserGroupSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupSearchLogic {
	return &UserGroupSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserGroupSearchLogic) UserGroupSearch(in *group_rpc.UserGroupSearchRequest) (*group_rpc.UserGroupSearchResponse, error) {
	// todo: add your logic here and delete this line

	return &group_rpc.UserGroupSearchResponse{}, nil
}
