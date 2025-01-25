package userlogic

import (
	"context"
	"fim_server/utils/stores/logs"
	"strconv"

	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"

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
		value, err := strconv.Atoi(key)
		if err != nil {
			logs.Error("转换失败", err.Error())
			continue
		}
		resp.UserIdList = append(resp.UserIdList, uint32(value))
	}

	return resp, nil
}
