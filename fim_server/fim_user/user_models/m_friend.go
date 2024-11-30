package user_models

import (
	"fim_server/common/models"
	"gorm.io/gorm"
)

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserId        uint      `json:"send_user_id"` // 发送人
	SendUserModel     UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId     uint      `json:"receive_user_id"` // 接收人
	ReceiveUserModel  UserModel `gorm:"foreignKey:ReceiveUserId" json:"-"`
	SendUserNotice    string    `gorm:"size:128" json:"send_user_notice"`    // 发起方备注
	ReceiveUserNotice string    `gorm:"size:128" json:"receive_user_notice"` // 接收方备注
}

// FriendAuthModel 好友验证表
type FriendAuthModel struct {
	models.Model
	SendUserId       uint                 `json:"send_user_id"` // 发送人
	SendUserModel    UserModel            `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId    uint                 `json:"receive_user_id"` // 接收人
	ReceiveUserModel UserModel            `gorm:"foreignKey:ReceiveUserId" json:"-"`
	Status           int8                 `json:"status"`                       // 0 等待验证  1 同意  2 拒绝  3 忽略  4 删除
	AuthMessage      string               `gorm:"size:128" json:"auth_message"` // 验证消息
	AuthQuestion     *models.AuthQuestion `json:"auth_question"`                // 验证问题
}

func (f *FriendModel) IsFriend(db *gorm.DB, userId uint, friendId uint) bool {
	err := db.Take(&f, "(send_user_id = ? and receive_user_id = ?) or (send_user_id = ? and receive_user_id = ?)",
		userId, friendId, friendId, userId).Error
	return err == nil
}

func (f *FriendModel) MeFriend(db *gorm.DB, userId uint) (list []FriendModel) {
	db.Find(&list, "send_user_id = ? or receive_user_id = ?", userId, userId)
	return
}

func (f *FriendModel) GetUserNotice(userId uint) string {
	if userId == f.SendUserId {
		return f.SendUserNotice
	}
	if userId == f.ReceiveUserId {
		return f.ReceiveUserNotice
	}
	return ""
}
