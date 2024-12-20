package group_models

import (
	"fim_server/models"
)

type GroupModel struct {
	models.Model
	Name               string             `gorm:"size:32" json:"name"`         // 群名
	Sign               string             `gorm:"size:128" json:"sign"`        // 简介
	Avatar             string             `gorm:"size:256" json:"avatar"`      // 群头像
	Leader             uint               `json:"leader"`                      // 群主
	Verify             int8               `gorm:"size:32" json:"verify"`       // 验证
	VerifyInfo         models.VerifyInfo  `json:"verify_info"`                 // 验证问题
	IsSearch           bool               `json:"is_search"`                   // is搜索
	IsInvite           bool               `json:"is_invite"`                   // is邀请
	IsTemporarySession bool               `json:"is_temporary_session"`        // is临时会话
	IsForbiddenSpeech  bool               `json:"is_forbidden_speech"`         // is禁言
	Size               int                `json:"size"`                        // 群规模 20 100 200 1000
	MemberList         []GroupMemberModel `gorm:"foreignKey:GroupId" json:"-"` // 群成员列表
}
