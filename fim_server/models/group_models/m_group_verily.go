package group_models

import (
	"fim_server/models"
)

// GroupValidModel 群验证表
type GroupValidModel struct {
	models.Model
	GroupId    uint              `json:"group_id"`              // 群ID
	UserId     uint              `json:"user_id"`               // 用户加群 退群
	Status     int8              `json:"status"`                // 状态
	Valid     int8              `gorm:"size:32" json:"valid"` // 验证
	ValidInfo models.ValidInfo `json:"valid_info"`           // 验证问题
	Type       int8              `json:"type"`                  // 1 加群 2 退群
	GroupModel GroupModel        `gorm:"foreignKey:GroupId" json:"-"`
}
