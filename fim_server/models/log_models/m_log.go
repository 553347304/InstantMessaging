package log_model

import "fim_server/models"

type LogModel struct {
	models.Model
	Type    int8   `json:"type"` // 日志类型  2 操作日志 3 运行日志
	IP      string `gorm:"size:32" json:"ip"`
	Addr    string `gorm:"size:64" json:"addr"`
	UserID  uint   `json:"userID"`
	Name    string `gorm:"size:64" json:"name"`
	Avatar  string `gorm:"size:256" json:"avatar"`
	Level   string `gorm:"size:12" json:"level"`
	Title   string `gorm:"size:32" json:"title"`
	Content string `json:"content"`                // 日志详情
	Service string `gorm:"size:32" json:"service"` // 服务  记录微服务的名称
	IsRead  bool   `json:"isRead"`
}
