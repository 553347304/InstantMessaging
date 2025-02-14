package userlogic

import (
	"context"
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/stores/conv"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineListLogic {
	return &UserOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOnlineListLogic) UserOnlineList(in *user_rpc.Empty) (*user_rpc.UserOnlineListResponse, error) {
	// todo: add your logic here and delete this line

	resp := new(user_rpc.UserOnlineListResponse)
	onlineMap := l.svcCtx.Redis.HGetAll("user_online").Val()
	for key, _ := range onlineMap {
		u, err := conv.Type(key).Uint64()
		if err != nil {
			continue
		}
		resp.UserIdList = append(resp.UserIdList, u)
	}

	return resp, nil
}
