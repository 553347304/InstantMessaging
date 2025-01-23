package friendlogic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"
	
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFriendLogic) IsFriend(in *user_rpc.IsFriendRequest) (*user_rpc.Empty, error) {
	// todo: add your logic here and delete this line
	
	if in.User1 == in.User2 {
		logs.Info("好友是自己")
		return nil, nil
	}
	
	var friend user_models.FriendModel
	err := l.svcCtx.DB.Take(&friend, "(send_user_id = ? and receive_user_id = ?) or (send_user_id = ? and receive_user_id = ?)",
		in.User1, in.User2, in.User2, in.User1).Error
	
	if err != nil {
		return nil, logs.Error("不是好友")
	}
	
	return &user_rpc.Empty{}, nil
}
