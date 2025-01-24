package admin

import (
	"context"
	"fim_server/models/setting_models"
	"fim_server/utils/stores/logs"
	
	"fim_server/service/api/setting/internal/svc"
	"fim_server/service/api/setting/internal/types"
)

type SettingInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettingInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingInfoUpdateLogic {
	return &SettingInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SettingInfoUpdateLogic) SettingInfoUpdate(req *setting_models.ConfigModel) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	
	var settingModel setting_models.ConfigModel
	l.svcCtx.DB.First(&settingModel)
	err = l.svcCtx.DB.Model(&settingModel).Updates(req).Error
	logs.Info("更新配置", err == nil)
	return
}
