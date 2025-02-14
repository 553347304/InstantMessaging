package user_models

import (
	"database/sql/driver"
	"encoding/json"
	"fim_server/models"
	"fim_server/models/mgorm"
)

// UserModel 用户表
type UserModel struct {
	models.Model
	Username        string          `gorm:"size:32" json:"username"` // 昵称
	Password        string          `gorm:"size:64" json:"password"` // 密码
	Sign            string          `gorm:"size:128" json:"sign"`    // 签名
	Avatar          string          `gorm:"size:256" json:"avatar"`  // 头像
	Ip              string          `gorm:"size:32" json:"ip"`       // IP
	Addr            string          `gorm:"size:64" json:"addr"`     // 地址
	Role            int32           `gorm:"size:32" json:"role"`     // 角色 1管理员 2普通用户
	OpenId          string          `gorm:"size:64" json:"open_id"`
	RegisterSource  string          `gorm:"size:16" json:"register_source"` // 注册来源 1手机号 2邮箱 3第三方
	UserConfigModel *UserConfigModel `gorm:"foreignKey:UserId" json:"user_config_model"`
	Top             TopModel        `json:"top"`
}
type TopModel struct {
	GroupId mgorm.Uint64 `json:"group_id"`
}

func (v TopModel) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *TopModel) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }
