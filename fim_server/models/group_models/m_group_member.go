package group_models

import (
	"fim_server/models"
)

// GroupMemberModel 群成员表
type GroupMemberModel struct {
	models.Model
	GroupId         uint       `json:"groupId"` // 群ID
	GroupModel      GroupModel `gorm:"foreignKey:GroupId" json:"-"`
	UserId          uint       `json:"user_id"`                    // 用户ID
	MemberName      string     `gorm:"size:32" json:"member_name"` // 群名称
	Role            int8        `json:"role"`                       // 1 群主 2 管理员 3 普通成员
	ForbiddenSpeech *int       `json:"forbidden_speech"`           // 禁言时间 单位分钟
}
