package logic

import (
	"context"
	"fim_server/models/chat_models"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/logs"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTopLogic {
	return &UserTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTopLogic) UserTop(req *types.UserTopRequest) (resp *types.UserTopResponse, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.UserRpc.Friend.IsFriend(l.ctx, &user_rpc.IsFriendRequest{User1: uint32(req.UserId), User2: uint32(req.FriendId)})
	if err != nil {
		return nil, logs.Error("不是好友")
	}

	var topUser chat_models.TopUserModel
	err1 := l.svcCtx.DB.Take(&topUser, "user_id = ? and top_user_id = ? or user_id = top_user_id", req.UserId, req.FriendId).Error
	if err1 != nil {
		// 没有置顶
		l.svcCtx.DB.Create(&chat_models.TopUserModel{
			UserId:    req.UserId,
			TopUserId: req.FriendId,
		})
		return
	}
	logs.Info(topUser)
	l.svcCtx.DB.Delete(&topUser)
	return
}
