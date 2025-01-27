package grouplogic

import (
	"context"

	"fim_server/service/rpc/group/group_rpc"
	"fim_server/service/rpc/group/internal/svc"
)

type IsInGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsInGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsInGroupMemberLogic {
	return &IsInGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsInGroupMemberLogic) IsInGroupMember(in *group_rpc.IsInGroupMemberRequest) (*group_rpc.EmptyResponse, error) {
	// todo: add your logic here and delete this line

	return &group_rpc.EmptyResponse{}, nil
}
