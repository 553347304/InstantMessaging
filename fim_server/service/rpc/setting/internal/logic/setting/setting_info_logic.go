package settinglogic

import (
	"context"
	"fim_server/models/setting_models"
	"fim_server/utils/stores/conv"
	
	"fim_server/service/rpc/setting/internal/svc"
	"fim_server/service/rpc/setting/setting_rpc"
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

func (l *SettingInfoLogic) SettingInfo(in *setting_rpc.Empty) (*setting_rpc.SettingInfoResponse, error) {
	// todo: add your logic here and delete this line
	
	resp := new(setting_rpc.SettingInfoResponse)
	var setting setting_models.ConfigModel
	if l.svcCtx.DB.First(&setting).Error == nil {
		resp.Data = conv.Json().Marshal(setting)
		return resp, nil
	}
	
	go l.svcCtx.DB.Create(&setting_models.SystemSetting)
	resp.Data = conv.Json().Marshal(setting_models.SystemSetting)
	return resp, nil
}
