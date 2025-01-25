package logic

import (
	"context"
	"fim_server/models/setting_models"
	"fim_server/service/api/setting/internal/svc"
	"fim_server/service/api/setting/internal/types"
	"fim_server/service/rpc/setting/setting_rpc"
	"fim_server/utils/stores/conv"
)

type SettingInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettingInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingInfoLogic {
	return &SettingInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SettingInfoLogic) SettingInfo(req *types.Empty) (resp *setting_models.ConfigModel, err error) {
	// todo: add your logic here and delete this line
	resp = new(setting_models.ConfigModel)
	info, err := l.svcCtx.SettingRpc.SettingInfo(l.ctx, &setting_rpc.Empty{})
	if err != nil {
		return nil, err
	}
	conv.Json().Unmarshal(info.Data, &resp)
	return resp, nil
}
