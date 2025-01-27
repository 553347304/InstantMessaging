package logic

import (
	"context"
	
	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"
)

type GroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupChatLogic {
	return &GroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupChatLogic) GroupChat(req *types.GroupChatRequest) (resp *types.GroupChatResponse, err error) {
	// todo: add your logic here and delete this line

	
	
	
	
	
	return
}
