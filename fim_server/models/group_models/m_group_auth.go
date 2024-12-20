package group_models

import (
	"fim_server/models"
)

// GroupAuthModel 群验证表
type GroupAuthModel struct {
	models.Model
	GroupId      uint                 `json:"group_id"`      // 群ID
	UserId       uint                 `json:"user_id"`       // 用户加群 退群
	Status       int8                 `json:"status"`        // 状态
	Verify             int8                  `gorm:"size:32" json:"verify"`       // 验证
	VerifyInfo         models.VerifyInfo  `json:"verify_info"`                 // 验证问题
	Type         int8                 `json:"type"`          // 1 加群 2 退群
}
