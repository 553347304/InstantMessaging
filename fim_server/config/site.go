package config

import (
	"fim_server/models/setting_models"
	"fim_server/utils/stores/method"
)

var SystemSetting = setting_models.ConfigModel{
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
	OpenLoginQQ: setting_models.OpenLoginQQ{
		Enable:   true,
		AppID:    "102550927",
		Key:      "rdEbkhT2RgovviQ0",
		Redirect: "http://tcbyj.cn/login/qq",
		WebPath:  "",
	},
}
