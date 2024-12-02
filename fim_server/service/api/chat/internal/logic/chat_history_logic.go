package logic

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/utils/src/sqls"
	"fim_server/utils/stores"
	"fim_server/utils/stores/method"

	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryLogic {
	return &ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *types.ChatHistoryResponse, err error) {
	// todo: add your logic here and delete this line

	chatList, total := sqls.GetList(chat_models.ChatModel{ReceiveUserId: req.UserId}, sqls.Mysql{
		DB:      l.svcCtx.DB.Where("send_user_id = ? or receive_user_id = ?", req.UserId, req.UserId),
		Preload: []string{"ReceiveUserModel", "SendUserModel.UserConfigModel"},
		PageInfo: stores.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	var userIdList []uint
	for _, model := range chatList {
		userIdList = append(userIdList, model.SendUserId)
		userIdList = append(userIdList, model.ReceiveUserId)
	}

	userIdList = method.Deduplication(userIdList)

	var list = make([]types.ChatHistory, 0)
	for _, model := range chatList {
		list = append(list, types.ChatHistory{
			ID:            model.ID,
			CreatedAt:     model.CreatedAt.String(),
			Message:       model.Message,
			SystemMessage: model.SystemMessage,
		})
	}
	resp = &types.ChatHistoryResponse{List: list, Total: int(total)}

	return
}
