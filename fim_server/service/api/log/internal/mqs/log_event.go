package mqs

import (
	"context"
	"fim_server/service/api/log/internal/svc"
)

type LogEvent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *LogEvent {
	return &LogEvent{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogEvent) Consume(ctx context.Context, key, val string) error {
	return nil
}
