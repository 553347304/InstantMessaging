package zrpc_client

import (
	"fim_server/common/zero_middleware"
	"fim_server/service/rpc/chat/chat_rpc"
	"fim_server/service/rpc/chat/client/chat"
	"fim_server/service/rpc/file/file"
	"fim_server/service/rpc/file/file_rpc"
	"fim_server/service/rpc/setting/client/setting"
	"fim_server/service/rpc/setting/setting_rpc"
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
type SettingRpc setting_rpc.SettingClient
type ChatRpc chat_rpc.ChatClient
type FileRpc file_rpc.FileClient

type rpcService struct{ rpc zrpc.RpcClientConf }

func Service(rpc zrpc.RpcClientConf) *rpcService { return &rpcService{rpc: rpc} }

func (r *rpcService) UserClient() UserRpc {
	return UserRpc{
		User:    user.NewUser(zrpc.MustNewClient(r.rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		Friend:  friend.NewFriend(zrpc.MustNewClient(r.rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
		Curtail: usercurtail.NewUserCurtail(zrpc.MustNewClient(r.rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor))),
	}
}
func (r *rpcService) SettingRpc() SettingRpc {
	return setting.NewSetting(zrpc.MustNewClient(r.rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor)))
}
func (r *rpcService) ChatRpc() ChatRpc {
	return chat.NewChat(zrpc.MustNewClient(r.rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor)))
}
func (r *rpcService) FileRpc() FileRpc {
	return file.NewFile(zrpc.MustNewClient(r.rpc, zrpc.WithUnaryClientInterceptor(zero_middleware.ClientInterceptor)))
}
