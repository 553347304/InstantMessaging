package logic

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method/method_list"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatSessionLogic {
	return &ChatSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatSessionLogic) ChatSession(req *types.ChatSessionRequest) (resp *types.ChatSessionResponse, err error) {
	// todo: add your logic here and delete this line

	type Data struct {
		SU         uint   `gorm:"column:s_u"`
		RU         uint   `gorm:"column:r_u"`
		MaxDate    string `gorm:"column:max_date"`
		MaxPreview string `gorm:"column:max_preview"`
		IsTop      bool   `gorm:"column:is_top"`
	}

	var chatList []Data
	chatResponse := sqls.GetListGroup(chat_models.ChatModel{}, sqls.Mysql{
		DB: l.svcCtx.DB.
			Select("least(send_user_id, receive_user_id) as s_u,"+
				"greatest(send_user_id, receive_user_id) as r_u,"+
				"count(id) as ct,"+
				"max(created_at) as max_date,"+
				"(select message_preview from chat_models "+
				"where (send_user_id = s_u and receive_user_id = r_u)"+
				"or (send_user_id = r_u and receive_user_id = s_u)"+
				"and id not in (select chat_id from user_chat_delete_models where user_id = ?)"+
				"order by created_at desc limit 1) as max_preview,"+
				"if((select 1 from top_user_models where user_id = ? "+
				"and (top_user_id = s_u or top_user_id = r_u)), 1, 0) as is_top",
				req.UserId, req.UserId).
			Where("send_user_id = ? or receive_user_id = ? and id not in (select chat_id from user_chat_delete_models where user_id = ?)",
				req.UserId, req.UserId, req.UserId).
			Group("least(send_user_id, receive_user_id)").
			Group("greatest(send_user_id, receive_user_id)"),
		PageInfo: src.PageInfo{
			Sort:  "is_top desc, max_date desc",
			Page:  req.Page,
			Limit: req.Limit,
		},
	},&chatList)

	var userIdList []uint32
	for _, data := range chatList {
		if data.RU != req.UserId {
			userIdList = append(userIdList, uint32(data.RU))
		}
		if data.SU != req.UserId {
			userIdList = append(userIdList, uint32(data.SU))
		}
		// 自己和自己聊
		if data.SU == req.UserId && req.UserId == data.RU {
			userIdList = append(userIdList, uint32(req.UserId))
		}
	}
	userIdList = method_list.Unique(userIdList) // 去重
	// 调用户服务
	response, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIdList,
	})
	if err != nil {
		return nil, logs.Error("用户服务错误")
	}

	var list = make([]types.ChatSession, 0)
	for _, data := range chatList {
		s := types.ChatSession{
			CreatedAt:      data.MaxDate,
			MessagePreview: data.MaxPreview,
			IsTop:          data.IsTop,
		}
		if data.RU != req.UserId {
			s.UserId = data.RU
		}
		if data.SU != req.UserId {
			s.UserId = data.SU
		}
		if data.SU == req.UserId && req.UserId == data.RU {
			s.UserId = data.SU
		}
		s.Avatar = response.UserInfo[uint32(s.UserId)].Avatar
		s.Name = response.UserInfo[uint32(s.UserId)].Name
		list = append(list, s)
	}
	return &types.ChatSessionResponse{List: list, Total: chatResponse.Total}, nil
}
