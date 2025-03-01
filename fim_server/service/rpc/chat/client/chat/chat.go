// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: chat.proto

package chat

import (
	"context"

	"fim_server/service/rpc/chat/chat_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChatTotalMessage          = chat_rpc.ChatTotalMessage
	UserListChatTotalRequest  = chat_rpc.UserListChatTotalRequest
	UserListChatTotalResponse = chat_rpc.UserListChatTotalResponse

	Chat interface {
		UserListChatTotal(ctx context.Context, in *UserListChatTotalRequest, opts ...grpc.CallOption) (*UserListChatTotalResponse, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) UserListChatTotal(ctx context.Context, in *UserListChatTotalRequest, opts ...grpc.CallOption) (*UserListChatTotalResponse, error) {
	client := chat_rpc.NewChatClient(m.cli.Conn())
	return client.UserListChatTotal(ctx, in, opts...)
}
