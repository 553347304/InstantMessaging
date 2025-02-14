package chatlogic

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/utils/stores/logs"
	
	"fim_server/service/rpc/chat/chat_rpc"
	"fim_server/service/rpc/chat/internal/svc"
)

type UserListChatTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListChatTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListChatTotalLogic {
	return &UserListChatTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	UserId uint32 `gorm:"column:user_id"`
	Count  uint32 `gorm:"column:count"`
}

func (l *UserListChatTotalLogic) UserListChatTotal(in *chat_rpc.UserListChatTotalRequest) (resp *chat_rpc.UserListChatTotalResponse, err error) {
	// todo: add your logic here and delete this line
	var sendScan []Data
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("send_user_id in ?", in.UserIDList).
		Group("send_user_id").
		Select("send_user_id as user_id", "count(id) as count").Scan(&sendScan)
	
	var receiveScan []Data
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("send_user_id in ?", in.UserIDList).
		Group("send_user_id").
		Select("send_user_id as user_id", "count(id) as count").Scan(&receiveScan)
	
	resp = new(chat_rpc.UserListChatTotalResponse)
	resp.Result = map[uint32]*chat_rpc.ChatTotalMessage{}
	
	for _, data := range sendScan {
		result, ok := resp.Result[data.UserId]
		if !ok {
			resp.Result[data.UserId] = &chat_rpc.ChatTotalMessage{
				SendMessageTotal: int32(data.Count),
			}
		} else {
			result.ReceiveMessageTotal = int32(data.Count)
		}
	}
	logs.Info(resp)
	return &chat_rpc.UserListChatTotalResponse{}, nil
}
