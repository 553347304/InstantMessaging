package log_models

import (
	"fim_server/models"
)

type LogModel struct {
	models.Model
	UserId  uint   `json:"user_id"`
	Name    string `gorm:"size:64" json:"name"`
	Avatar  string `gorm:"size:256" json:"avatar"`
	IP      string `gorm:"size:32" json:"ip"`
	Addr    string `gorm:"size:64" json:"addr"`
	Service string `gorm:"size:32" json:"service"` // 服务  记录微服务的名称
	Type    string `json:"type"`                   // 日志类型  操作日志 | 运行日志
	Level   string `gorm:"size:12" json:"level"`
	Title   string `gorm:"size:32" json:"title"`
	Content string `json:"content"` // 日志详情
	IsRead  bool   `json:"isRead"`
}

