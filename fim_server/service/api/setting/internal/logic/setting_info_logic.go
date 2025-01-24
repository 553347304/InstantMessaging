package logic

import (
	"context"
	"fim_server/models/setting_models"
	"fim_server/utils/stores/method"
	
	"fim_server/service/api/setting/internal/svc"
	"fim_server/service/api/setting/internal/types"
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
	
	var setting setting_models.ConfigModel
	if l.svcCtx.DB.First(&setting).Error == nil {
		return &setting, nil
	}
	init := setting_models.ConfigModel{
		Site: setting_models.Site{
			CreatedAt:   method.Time().NowDay,
			BeiAn:       "津ICP备2024017367号-1",
			Version:     "1.0.0",
			ImageQQ:     "https://vip.123pan.cn/1821560246/website/image/code/QQ%E4%BA%8C%E7%BB%B4%E7%A0%81.jpg",
			ImageWechat: "https://vip.123pan.cn/1821560246/website/image/code/%E5%BE%AE%E4%BF%A1%E4%BA%8C%E7%BB%B4%E7%A0%81.jpg",
			UrlBiliBili: "https://space.bilibili.com/59452692",
			UrlGitee:    "https://gitee.com/baiyins",
			UrlGithub:   "https://github.com/553347304",
		},
	}
	go l.svcCtx.DB.Create(&init)
	
	return &init, nil
}
