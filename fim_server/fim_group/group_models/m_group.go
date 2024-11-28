package group_models

import (
	"fim_server/common/models"
)

type Group struct {
	models.Model
	Name               string               `gorm:"size:32" json:"name"`         // 群名
	Sign               string               `gorm:"size:128" json:"sign"`        // 简介
	Avatar             string               `gorm:"size:256" json:"avatar"`      // 群头像
	Leader             string               `json:"leader"`                      // 群主
	AuthMessage        int8                 `gorm:"size:32" json:"auth_message"` // 群验证
	AuthQuestion       *models.AuthQuestion `json:"auth_question"`               // 群验证问题
	IsSearch           string               `json:"is_search"`                   // is搜索
	IsInvite           bool                 `json:"is_invite"`                   // is邀请
	IsTemporarySession bool                 `json:"is_temporary_session"`        // is临时会话
	IsForbiddenSpeech  bool                 `json:"is_forbidden_speech"`         // is禁言
	Size               int                  `json:"size"`                        // 群规模 20 100 200 1000
}
