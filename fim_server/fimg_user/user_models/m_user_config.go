package user_models

import (
	"fim_server/common/models"
)

// UserConfig 用户配置表
type UserConfig struct {
	UserId        uint    `json:"UserId"`
	UserModel     User    `gorm:"foreignKey:UserId" json:"-"`
	RecallMessage *string `gorm:"size:32" json:"RecallMessage"` // 撤回消息内容
	FriendOnline  bool    `json:"FriendOnline"`                 // 好友上线
	Sound         bool    `json:"sound"`                        // 好友上线声音
	SecureLink    bool    `json:"SecureLink"`                   // 安全链接
	SavePassword  bool    `json:"savePassword"`                 // 保存密码

	// 防骚扰
	SearchUser   int8                 `json:"SearchUser"`   // 别人查找到你的方式
	Auth         int8                 `json:"auth"`         // 好友验证
	AuthQuestion *models.AuthQuestion `json:"authQuestion"` // 验证问题
}
