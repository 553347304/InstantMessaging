package group_models

import (
	"fim_server/models"
	"fim_server/models/mgorm"
	"fim_server/models/mtype"
)

// GroupMessageModel 群聊消息表
type GroupMessageModel struct {
	models.Model
	GroupId      uint64           `json:"groupId"` // 发送人
	GroupModel   GroupModel       `gorm:"foreignKey:GroupId" json:"-"`
	SendUserId   uint64           `json:"send_user_id"`
	MemberId     uint64           `json:"member_id"`                    // 群成员ID
	MemberModel  GroupMemberModel `gorm:"foreignKey:MemberId" json:"-"` // 对应的群成员
	Type         mtype.Int8       `json:"type"`                         // 消息类型
	Preview      string           `gorm:"size:64" json:"preview"`       // 消息预览
	Message      mtype.Message    `json:"message"`                      // 消息内容
	DeleteUserId mgorm.Uint64     `json:"delete_user_id"`               // 用户删除的聊天记录
}
