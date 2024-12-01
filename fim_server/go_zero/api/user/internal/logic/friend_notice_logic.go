package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/go_zero/api/user/internal/svc"
	"fim_server/go_zero/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendNoticeLogic {
	return &FriendNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendNoticeLogic) FriendNotice(req *types.FriendNoticeUpdateRequest) (resp *types.FriendNoticeUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	var friend user_models.FriendModel
	if !friend.IsFriend(l.svcCtx.DB, req.UserId, req.FriendId) {
		return nil, logs.Error("不是好友")
	}

	if req.UserId == friend.SendUserId {
		if friend.SendUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("send_user_notice", req.Notice)
	}
	if req.UserId == friend.ReceiveUserId {
		if friend.ReceiveUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("receive_user_notice", req.Notice)
	}

	return
}
