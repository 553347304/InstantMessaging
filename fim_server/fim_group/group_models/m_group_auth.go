package group_models

import (
	"fim_server/common/models"
)

// GroupAuth 群验证表
type GroupAuth struct {
	models.Model
	GroupId      uint                 `json:"groupId"`      // 群ID
	UserId       uint                 `json:"userId"`       // 用户加群 退群
	Status       int8                 `json:"status"`       // 状态
	Auth         int8                 `json:"auth"`         // 群验证
	AuthQuestion *models.AuthQuestion `json:"authQuestion"` // 群验证问题
	Type         int8                 `json:"type"`         // 1 加群 2 退群
}
