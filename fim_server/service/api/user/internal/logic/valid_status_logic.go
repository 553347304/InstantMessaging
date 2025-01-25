package logic

import (
	"context"
	"fim_server/models/mtype"
	"fim_server/models/user_models"
	"fim_server/service/rpc/chat/chat"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidStatusLogic {
	return &ValidStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidStatusLogic) ValidStatus(req *types.ValidStatusRequest) (resp *types.ValidStatusResponse, err error) {
	// todo: add your logic here and delete this line

	var friendVerify user_models.FriendValidModel
	err = l.svcCtx.DB.Take(&friendVerify, "id = ? and receive_user_id = ?", req.ValidId, req.UserId).Error

	if err != nil {
		return nil, logs.Error("验证记录不存在")
	}
	if friendVerify.Status != 0 {
		return nil, logs.Error("不可更改状态")
	}

	switch req.Status {
	case 1:
		// 同意
		friendVerify.ReceiveStatus = 1
		l.svcCtx.DB.Create(&user_models.FriendModel{
			SendUserId:    friendVerify.SendUserId,
			ReceiveUserId: friendVerify.ReceiveUserId,
		})

		message := mtype.MessageArray{
			{Type: mtype.MessageType.Text, Content: "已添加你为好友"},
		}
		byteData := conv.Json().Marshal(message)

		// 给对方发消息
		_, err = l.svcCtx.ChatRpc.UserChat(l.ctx, &chat.UserChatRequest{
			SendUserId:    uint32(req.UserId),
			ReceiveUserId: uint32(req.ValidId),
			Message:       byteData,
			SystemMessage: nil,
		})
		if err != nil {
			logs.Error("发送消息失败", err)
		}
	case 2:
		friendVerify.ReceiveStatus = 2 // 拒绝
	case 3:
		friendVerify.ReceiveStatus = 3 // 忽略
	case 4:
		friendVerify.Status = 4 // 删除
		l.svcCtx.DB.Delete(&friendVerify)
		return nil, nil
	}
	friendVerify.Status = friendVerify.ReceiveStatus
	logs.Info(friendVerify.Status)
	l.svcCtx.DB.Model(&friendVerify).Updates(map[string]any{
		"status":         friendVerify.Status,
		"receive_status": friendVerify.ReceiveStatus,
	})
	return
}
