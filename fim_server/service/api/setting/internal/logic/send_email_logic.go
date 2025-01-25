package logic

import (
	"context"
	"fim_server/utils/open_api/open_api_qq"
	"fim_server/utils/stores/conv"
	
	"fim_server/service/api/setting/internal/svc"
	"fim_server/service/api/setting/internal/types"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.SendEmailResponse) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	
	go open_api_qq.SendEmail(conv.Struct(open_api_qq.EmailConfig{}).Type(&req))
	
	return
}
