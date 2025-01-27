// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package friend

import (
	"context"

	"fim_server/service/rpc/user/user_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CurtailResponse        = user_rpc.CurtailResponse
	Empty                  = user_rpc.Empty
	FriendListResponse     = user_rpc.FriendListResponse
	ID                     = user_rpc.ID
	IdList                 = user_rpc.IdList
	IsFriendRequest        = user_rpc.IsFriendRequest
	UserCreateRequest      = user_rpc.UserCreateRequest
	UserCreateResponse     = user_rpc.UserCreateResponse
	UserInfo               = user_rpc.UserInfo
	UserInfoResponse       = user_rpc.UserInfoResponse
	UserOnlineListResponse = user_rpc.UserOnlineListResponse

	Friend interface {
		IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*Empty, error)
		FriendList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FriendListResponse, error)
	}

	defaultFriend struct {
		cli zrpc.Client
	}
)

func NewFriend(cli zrpc.Client) Friend {
	return &defaultFriend{
		cli: cli,
	}
}

func (m *defaultFriend) IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user_rpc.NewFriendClient(m.cli.Conn())
	return client.IsFriend(ctx, in, opts...)
}

func (m *defaultFriend) FriendList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FriendListResponse, error) {
	client := user_rpc.NewFriendClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}
