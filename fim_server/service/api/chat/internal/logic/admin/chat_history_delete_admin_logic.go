package admin

import (
	"context"
	"fim_server/models/chat_models"
	
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
)

type ChatHistoryDeleteAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryDeleteAdminLogic {
	return &ChatHistoryDeleteAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryDeleteAdminLogic) ChatHistoryDeleteAdmin(req *types.RequestDelete) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	
	var messageList []chat_models.ChatModel
	l.svcCtx.DB.Find(&messageList, "id in ?", req.IdList)
	return
}
