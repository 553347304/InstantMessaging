package user_models

import (
	"fim_server/common/models"
)

// Friend 好友表
type Friend struct {
	models.Model
	SendUserId       uint   `json:"sendUserId"` // 发送人
	SendUserModel    User   `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId    uint   `json:"receiveUserId"` // 接收人
	ReceiveUserModel User   `gorm:"foreignKey:ReceiveUserId" json:"-"`
	Notice           string `gorm:"size:128" json:"notice"` // 备注
}

// FriendAuth 好友验证表
type FriendAuth struct {
	models.Model
	SendUserId       uint                 `json:"sendUserId"` // 发送人
	SendUserModel    User                 `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId    uint                 `json:"receiveUserId"` // 接收人
	ReceiveUserModel User                 `gorm:"foreignKey:ReceiveUserId" json:"-"`
	Status           int8                 `json:"status"`                      // 0 等待验证  1 同意  2 拒绝  3 忽略
	AuthMessage      string               `gorm:"size:128" json:"authMessage"` // 验证消息
	AuthQuestion     *models.AuthQuestion `json:"authQuestion"`                // 验证问题
}
