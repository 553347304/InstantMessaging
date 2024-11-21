package group_models

import (
	"fim_server/common/models"
)

type Group struct {
	models.Model
	Name               string               `gorm:"size:32" json:"name"`        // 群名
	Abstract           string               `gorm:"size:128" json:"abstract"`   // 简介
	Avatar             string               `gorm:"size:256" json:"avatar"`     // 群头像
	Leader             string               `json:"leader"`                     // 群主
	AuthMessage        int8                 `gorm:"size:32" json:"authMessage"` // 群验证
	AuthQuestion       *models.AuthQuestion `json:"authQuestion"`               // 群验证问题
	IsSearch           string               `json:"isSearch"`                   // is搜索
	IsInvite           bool                 `json:"isInvite"`                   // is邀请
	IsTemporarySession bool                 `json:"isTemporarySession"`         // is临时会话
	IsForbiddenSpeech  bool                 `json:"isForbiddenSpeech"`          // is禁言
	Size               int                  `json:"size"`                       // 群规模 20 100 200 1000
}
