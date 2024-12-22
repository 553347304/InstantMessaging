package user_models

import (
	"fim_server/models"
)

// UserConfigModel 用户配置表
type UserConfigModel struct {
	models.Model
	UserId        uint      `json:"userId"`
	UserModel     UserModel `gorm:"foreignKey:UserId" json:"-"`
	RecallMessage *string   `gorm:"size:32" json:"recall_message"` // 撤回消息内容
	FriendOnline  bool      `json:"friend_online"`                 // 好友上线
	Sound         bool      `json:"sound"`                         // 好友上线声音
	SecureLink    bool      `json:"secure_link"`                   // 安全链接
	SavePassword  bool      `json:"save_password"`                 // 保存密码

	// 防骚扰
	SearchUser int8              `json:"search_user"` // 别人查找到你的方式
	Valid       int8              `json:"valid"`        // 好友验证
	ValidInfo models.ValidInfo `json:"valid_info"` // 验证问题
	Online     bool              `json:"online"`      // 是否在线
}
