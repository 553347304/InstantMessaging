package user_models

import (
	"fim_server/common/models"
)

// User 用户表
type User struct {
	models.Model
	Name           string      `gorm:"size:32" json:"name"`    // 昵称
	Password       string      `gorm:"size:64" json:"-"`       // 密码
	Sign           string      `gorm:"size:128" json:"sign"`   // 签名
	Avatar         string      `gorm:"size:256" json:"avatar"` // 头像
	IP             string      `gorm:"size:32" json:"ip"`      // IP
	Addr           string      `gorm:"size:64" json:"addr"`    // 地址
	Role           int8        `json:"role"`                   // 角色 1管理员 2普通用户
	OpenId         string      `gorm:"size:64" json:"-"`
	RegisterSource string      `gorm:"size:16" json:"registerSource"` // 注册来源 1手机号 2邮箱 3第三方
	UserConfig     *UserConfig `gorm:"foreignKey:UserId" json:"UserConfig"`
}
