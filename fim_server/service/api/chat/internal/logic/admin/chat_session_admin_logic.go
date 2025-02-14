package admin

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
)

type ChatSessionAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatSessionAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatSessionAdminLogic {
	return &ChatSessionAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatSessionAdminLogic) ChatSessionAdmin(req *types.ChatSessionAdminRequest) (resp *types.UserInfoListResponse, err error) {
	// todo: add your logic here and delete this line
	var sendUserIDList []uint64
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("receive_user_id = ?", req.ReceiveUserId).
		Group("send_user_id").
		Select("send_user_id").Scan(&sendUserIDList)

	userRpc, err := l.svcCtx.UserRpc.User.UserInfo(l.ctx, &user_rpc.IdList{Id: sendUserIDList})
	if err != nil {
		logs.Error(err)
	}

	resp = new(types.UserInfoListResponse)
	for u, info := range userRpc.InfoList {
		resp.List = append(resp.List, types.UserInfo{
			UserId: u,
			Avatar: info.Avatar,
			Username:   info.Username,
		})
	}
	resp.Total = userRpc.Total
	return
}
