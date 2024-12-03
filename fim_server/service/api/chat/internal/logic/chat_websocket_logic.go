package logic

import (
	"context"

	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatWebsocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatWebsocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWebsocketLogic {
	return &ChatWebsocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatWebsocketLogic) ChatWebsocket(req *types.ChatRequest) (resp *types.ChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
