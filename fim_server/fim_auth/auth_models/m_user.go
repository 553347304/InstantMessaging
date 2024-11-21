package models

import (
	"fim_server/common/models"
)

// User 用户表
type User struct {
	models.Model
	Name     string `gorm:"size:32" json:"name"`     // 昵称
	Password string `gorm:"size:64" json:"password"` // 密码
	Sign     string `gorm:"size:128" json:"sign"`    // 签名
	Avatar   string `gorm:"size:256" json:"avatar"`  // 头像
	IP       string `gorm:"size:32" json:"ip"`       // IP
	Addr     string `gorm:"size:64" json:"addr"`     // 地址
}
