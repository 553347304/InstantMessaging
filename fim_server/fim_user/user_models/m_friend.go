package user_models

import (
	"fim_server/common/models"
	"gorm.io/gorm"
)

// Friend 好友表
type Friend struct {
	models.Model
	SendUserId        uint   `json:"send_user_id"` // 发送人
	SendUserModel     User   `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId     uint   `json:"receive_user_id"` // 接收人
	ReceiveUserModel  User   `gorm:"foreignKey:ReceiveUserId" json:"-"`
	SendUserNotice    string `gorm:"size:128" json:"send_user_notice"`    // 发起方备注
	ReceiveUserNotice string `gorm:"size:128" json:"receive_user_notice"` // 接收方备注
}

// FriendAuth 好友验证表
type FriendAuth struct {
	models.Model
	SendUserId       uint                 `json:"send_user_id"` // 发送人
	SendUserModel    User                 `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId    uint                 `json:"receive_user_id"` // 接收人
	ReceiveUserModel User                 `gorm:"foreignKey:ReceiveUserId" json:"-"`
	Status           int8                 `json:"status"`                       // 0 等待验证  1 同意  2 拒绝  3 忽略
	AuthMessage      string               `gorm:"size:128" json:"auth_message"` // 验证消息
	AuthQuestion     *models.AuthQuestion `json:"auth_question"`                // 验证问题
}

func (f *Friend) IsFriend(db *gorm.DB, userId uint, friendId uint) bool {
	err := db.Take(&f, "(send_user_id = ? and receive_user_id = ?) or (send_user_id = ? and receive_user_id = ?)",
		userId, friendId, friendId, userId)
	return err != nil
}

func (f *Friend) GetUserNotice(userId uint) string {
	if userId == f.SendUserId {
		return f.SendUserNotice
	}
	if userId == f.ReceiveUserId {
		return f.ReceiveUserNotice
	}
	return ""
}
