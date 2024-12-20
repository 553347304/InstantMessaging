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
	UserChatRequest  = chat_rpc.UserChatRequest
	UserChatResponse = chat_rpc.UserChatResponse

	Chat interface {
		UserChat(ctx context.Context, in *UserChatRequest, opts ...grpc.CallOption) (*UserChatResponse, error)
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

func (m *defaultChat) UserChat(ctx context.Context, in *UserChatRequest, opts ...grpc.CallOption) (*UserChatResponse, error) {
	client := chat_rpc.NewChatClient(m.cli.Conn())
	return client.UserChat(ctx, in, opts...)
}
