package logic

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/models/mtype"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	
	"fim_server/service/rpc/chat/chat_rpc"
	"fim_server/service/rpc/chat/internal/svc"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatLogic {
	return &UserChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChatLogic) UserChat(in *chat_rpc.UserChatRequest) (*chat_rpc.UserChatResponse, error) {
	// todo: add your logic here and delete this line
	
	var message mtype.MessageArray
	conv.Unmarshal(in.Message, &message)
	var systemMessage *mtype.SystemMessage
	conv.Unmarshal(in.Message, &systemMessage)
	chat := chat_models.ChatModel{
		SendUserId:    uint(in.SendUserId),
		ReceiveUserId: uint(in.ReceiveUserId),
		Message:       message,
		SystemMessage: systemMessage,
	}
	chat.Preview = chat.PreviewMethod()
	
	err := l.svcCtx.DB.Create(&chat).Error
	if err != nil {
		return nil, logs.Error(err)
	}
	
	return &chat_rpc.UserChatResponse{}, nil
}
