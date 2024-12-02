package logic

import (
	"context"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"

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

	l.svcCtx.DB.Raw(`select *
from (select least(send_user_id, receive_user_id)    as s_u,
             greatest(send_user_id, receive_user_id) as r_u,
             count(id)                           as ct,
             max(created_at)                     as max_date,
             (select message_preview
              from chat_models
              where (send_user_id = s_u and receive_user_id = r_u)
                 or (send_user_id = r_u and receive_user_id = s_u)
              order by created_at desc limit 1)  as max_preview,
             if((select 1 from top_user_models where user_id = 1 and (top_user_id = s_u or top_user_id = r_u)), 1, 0) as is_top
      from chat_models
      where send_user_id = 1 or receive_user_id = 1
      group by least(send_user_id, receive_user_id), greatest(send_user_id, receive_user_id)) as subquery
order by is_top desc, max_date desc limit 10 offset 0;`).Find(&chatList)
	// total := sqls.GetListGroup(chat_models.ChatModel{}, &chatList, sqls.Mysql{
	// 	DB: l.svcCtx.DB.
	// 		Select("least(send_user_id, receive_user_id) as s_u").
	// 		Select("greatest(send_user_id, receive_user_id) as r_u").
	// 		Select("max(created_at) as max_date").
	// 		Select("(select message_preview").
	// 		Select("from chat_models").
	// 		Select("where (send_user_id = s_u and receive_user_id = r_u)").
	// 		Select("or (send_user_id = r_u and receive_user_id = s_u)").
	// 		Select("order by created_at desc limit 1)").
	// 		Select("if((select 1 from top_user_models where user_id = 1 and (op_user_id = s_u or top_user_id = r_u)), 1, 0) as isTop").
	// 		Where("send_user_id = ? or receive_user_id = ?", 1, req.UserId).
	// 		Group("least(send_user_id, receive_user_id)").
	// 		Group("greatest(send_user_id, receive_user_id)"),
	// 	PageInfo: src.PageInfo{
	// 		Sort:  "max_date desc",
	// 		Page:  req.Page,
	// 		Limit: req.Limit,
	// 	},
	// })

	var userIdList []uint32
	for _, data := range chatList {
		if data.RU != req.UserId {
			userIdList = append(userIdList, uint32(data.RU))
		}
		if data.SU != req.UserId {
			userIdList = append(userIdList, uint32(data.SU))
		}
	}
	userIdList = method.Deduplication(userIdList) // 去重
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
		s.Avatar = response.UserInfo[uint32(s.UserId)].Avatar
		s.Name = response.UserInfo[uint32(s.UserId)].Name
		list = append(list, s)
	}
	return &types.ChatSessionResponse{List: list, Total: int64(len(list))}, nil
}
