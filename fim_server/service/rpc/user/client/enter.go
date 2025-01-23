package client

import (
	"fim_server/common/zero_middleware"
	"fim_server/service/rpc/user/client/friend"
	"fim_server/service/rpc/user/client/user"
	"fim_server/service/rpc/user/client/usercurtail"
	"fim_server/service/rpc/user/user_rpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type UserRpc struct {
	User    user_rpc.UserClient
	Friend  user_rpc.FriendClient
	Curtail user_rpc.UserCurtailClient
}

func UserClient(rpc zrpc.RpcClientConf) UserRpc {
	return UserRpc{
		User:    user.NewUser(zrpc.MustNewClient(rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		Friend:  friend.NewFriend(zrpc.MustNewClient(rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		Curtail: usercurtail.NewUserCurtail(zrpc.MustNewClient(rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
	}
}
