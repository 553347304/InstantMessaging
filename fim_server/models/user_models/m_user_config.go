package user_models

import (
	"fim_server/models"
)




// UserConfigModel 用户配置表
type UserConfigModel struct {
	models.Model
	UserId        uint64    `json:"user_id"`
	UserModel     UserModel `gorm:"foreignKey:UserId" json:"-"`
	RecallMessage *string   `gorm:"size:32" json:"recall_message"` // 撤回消息内容
	FriendOnline  bool      `json:"friend_online"`                 // 好友上线
	Sound         bool      `json:"sound"`                         // 好友上线声音
	SecureLink    bool      `json:"secure_link"`                   // 安全链接
	SavePassword  bool      `json:"save_password"`                 // 保存密码
	
	// 防骚扰
	SearchUser int32            `json:"search_user"` // 别人查找到你的方式
	Valid      int32            `json:"valid"`       // 好友验证  0:禁止加我为好友 | 1:允许任何人添加 | 2:需要验证 | 3:需要正确回答问题 | 4:需要回答问题并由我确认
	ValidInfo  models.ValidInfo `json:"valid_info"`  // 验证问题
	Online     bool             `json:"online"`      // 是否在线
	
	// 限制
	CurtailChat        bool `json:"curtail_chat"`         // 限制聊天
	CurtailAddUser     bool `json:"curtail_add_user"`     // 限制加人
	CurtailCreateGroup bool `json:"curtail_create_group"` // 限制建群
	CurtailAddGroup    bool `json:"curtail_add_group"`    // 限制加群
}
