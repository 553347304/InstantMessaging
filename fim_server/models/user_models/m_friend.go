package user_models

import (
	"fim_server/models"
	"gorm.io/gorm"
)

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserId        uint64      `json:"send_user_id"` // 发送人
	SendUserModel     UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId     uint64      `json:"receive_user_id"` // 接收人
	ReceiveUserModel  UserModel `gorm:"foreignKey:ReceiveUserId" json:"-"`
	SendUserNotice    string    `gorm:"size:128" json:"send_user_notice"`    // 发起方备注
	ReceiveUserNotice string    `gorm:"size:128" json:"receive_user_notice"` // 接收方备注
}

// FriendValidModel 好友验证表
type FriendValidModel struct {
	models.Model
	SendUserId       uint64             `json:"send_user_id"` // 发送人
	SendUserModel    UserModel        `gorm:"foreignKey:SendUserId" json:"-"`
	ReceiveUserId    uint64             `json:"receive_user_id"` // 接收人
	ReceiveUserModel UserModel        `gorm:"foreignKey:ReceiveUserId" json:"-"`
	Status           int32             `json:"status"`                        // 0 等待验证  1 同意  2 拒绝  3 忽略  4 删除
	SendStatus       int32             `json:"send_status"`                   // 发送方状态
	ReceiveStatus    int32             `json:"receive_status"`                // 接收方状态
	ValidMessage     string           `gorm:"size:128" json:"valid_message"` // 验证消息
	ValidInfo        models.ValidInfo `json:"valid_info"`                    // 验证问题
}


func (f *FriendModel) IsFriend(db *gorm.DB, UserId uint64, friendId uint64) bool {
	err := db.Take(&f, "(send_user_id = ? and receive_user_id = ?) or (send_user_id = ? and receive_user_id = ?)",
		UserId, friendId, friendId, UserId).Error
	return err == nil
}

func (f *FriendModel) MeFriend(db *gorm.DB, UserId uint64) (list []FriendModel) {
	db.Find(&list, "send_user_id = ? or receive_user_id = ?", UserId, UserId)
	return
}

func (f *FriendModel) GetUserNotice(UserId uint64) string {
	if UserId == f.SendUserId {
		return f.SendUserNotice
	}
	if UserId == f.ReceiveUserId {
		return f.ReceiveUserNotice
	}
	return ""
}
