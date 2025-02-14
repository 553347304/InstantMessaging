package log_models

import (
	"fim_server/models"
)

type LogModel struct {
	models.Model
	UserId   uint64 `json:"user_id"`
	Username string `gorm:"size:64" json:"username"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	Ip       string `gorm:"size:32" json:"ip"`
	Addr     string `gorm:"size:64" json:"addr"`
	Service  string `gorm:"size:32" json:"service"` // 服务  记录微服务的名称
	Type     string `json:"type"`                   // 日志类型  操作日志 | 运行日志
	Level    string `gorm:"size:12" json:"level"`
	Content  string `json:"content"` // 日志详情
}
