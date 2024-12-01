package group_models

import (
	models2 "fim_server/models"
)

// GroupAuthModel 群验证表
type GroupAuthModel struct {
	models2.Model
	GroupId      uint                  `json:"group_id"`      // 群ID
	UserId       uint                  `json:"user_id"`       // 用户加群 退群
	Status       int8                  `json:"status"`        // 状态
	Auth         int8                  `json:"auth"`          // 群验证
	AuthQuestion *models2.AuthQuestion `json:"auth_question"` // 群验证问题
	Type         int8                  `json:"type"`          // 1 加群 2 退群
}
