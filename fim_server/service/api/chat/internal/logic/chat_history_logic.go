package logic

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/models/mtype"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"
	"fmt"
	
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

type ChatHistory struct {
	ID        uint64          `json:"id"`
	IsMe      bool          `json:"is_me"` // 哪条消息是我发的
	Type      mtype.Int8    `json:"type"`
	Preview   string        `json:"preview"`
	Message   mtype.Message `json:"message"`
	CreatedAt string        `json:"created_at"` // 消息时间
}
type ChatHistoryResponse struct {
	Total       int64          `json:"total"`
	SendUser    mtype.UserInfo `json:"sendUser"`
	ReceiveUser mtype.UserInfo `json:"receive_user"`
	List        []ChatHistory  `json:"list"`
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *ChatHistoryResponse, err error) {
	// todo: add your logic here and delete this line
	
	_, err = l.svcCtx.UserRpc.Friend.IsFriend(l.ctx, &user_rpc.IsFriendRequest{User1: uint32(req.UserId), User2: uint32(req.FriendId)})
	if err != nil {
		return nil, err
	}
	
	chatList := src.Mysql(src.ServiceMysql[chat_models.ChatModel]{
		DB: l.svcCtx.DB.Where("(send_user_id = ? and receive_user_id = ?) or (receive_user_id = ? and send_user_id = ?)"+
			" and delete_user_id not like ?", req.UserId, req.FriendId, req.UserId, req.FriendId,
			fmt.Sprintf("%%\"%d\"%%", req.UserId)),
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
	}).GetList()
	
	var UserIDList []uint64
	for _, model := range chatList.List {
		UserIDList = append(UserIDList, model.SendUserId)
		UserIDList = append(UserIDList, model.ReceiveUserId)
	}
	UserIDList = method.List(UserIDList).Unique() // 去重
	// 调用户服务
	userResponse, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: UserIDList})
	if err != nil {
		return nil, logs.Error("用户服务错误")
	}
	
	var list = make([]ChatHistory, 0)
	var sendUser, receiveUser mtype.UserInfo
	for i, model := range chatList.List {
		
		if i == 0 {
			sendUser = mtype.UserInfo{
				UserId:     model.SendUserId,
				Username:   userResponse.InfoList[model.SendUserId].Username,
				Avatar: userResponse.InfoList[model.SendUserId].Avatar,
			}
			receiveUser = mtype.UserInfo{
				UserId:     model.ReceiveUserId,
				Username:   userResponse.InfoList[model.ReceiveUserId].Username,
				Avatar: userResponse.InfoList[model.ReceiveUserId].Avatar,
			}
		}
		info := ChatHistory{
			ID:        model.ID,
			Preview:   model.Preview,
			Type:      model.Type,
			Message:   model.Message,
			CreatedAt: model.CreatedAt.String(),
		}
		if model.SendUserId == req.UserId {
			info.IsMe = true
		}
		list = append(list, info)
		
	}
	resp = &ChatHistoryResponse{
		Total:       chatList.Total,
		SendUser:    sendUser,
		ReceiveUser: receiveUser,
		List:        list,
	}
	
	return
}
