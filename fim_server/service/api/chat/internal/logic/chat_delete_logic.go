package logic

import (
	"context"
	
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatDeleteLogic {
	return &ChatDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatDeleteLogic) ChatDelete(req *types.ChatDeleteRequest) (resp *types.ChatDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	// var chatList []chat_models.ChatModel
	// l.svcCtx.DB.Find(&chatList, req.IdList)
	// 
	// var userDeleteChatList []chat_models.UserChatDeleteModels
	// l.svcCtx.DB.Find(&userDeleteChatList, req.IdList)
	// chatDeleteMap := map[uint]struct{}{}
	// for _, model := range userDeleteChatList {
	// 	chatDeleteMap[model.ChatId] = struct{}{}
	// }
	// 
	// var deleteChatIdList []chat_models.UserChatDeleteModels
	// if len(chatList) > 0 {
	// 	for _, model := range chatList {
	// 		if !(model.SendUserID == req.UserID || model.ReceiveUserID == req.UserID) {
	// 			logs.Info("不是好友")
	// 			continue
	// 		}
	// 		_, ok := chatDeleteMap[model.ID]
	// 		if ok {
	// 			logs.Info("已经删除过了")
	// 			continue
	// 		}
	// 		deleteChatIdList = append(deleteChatIdList, chat_models.UserChatDeleteModels{
	// 			UserID: req.UserID,
	// 			ChatId: model.ID,
	// 		})
	// 	}
	// }
	// if len(deleteChatIdList) > 0 {
	// 	l.svcCtx.DB.Create(&deleteChatIdList)
	// }

	// logs.Info("删除聊天记录条数", len(deleteChatIdList))
	return
}
