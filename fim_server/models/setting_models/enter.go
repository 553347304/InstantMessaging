package setting_models

import (
	"database/sql/driver"
	"encoding/json"
)

type Site struct {
	CreatedAt   string `json:"created_at"` // 年月日
	BeiAn       string `json:"bei_an"`
	Version     string `json:"version"`
	ImageQQ     string `json:"image_qq"`
	ImageWechat string `json:"image_wechat"`
	UrlBiliBili string `json:"url_bili_bili"`
	UrlGitee    string `json:"url_gitee"`
	UrlGithub   string `json:"url_github"`
}
type QQ struct {
	Enable   bool   `json:"enable"` // 是否启用
	AppID    string `json:"app_id"`
	Key      string `json:"key"`
	Redirect string `json:"redirect"` // 登录之后的回调地址
	WebPath  string `json:"webPath"`  // 点击跳转的路径
}
type OpenLogin struct {
	QQ QQ `json:"qq"`
}
type ConfigModel struct {
	ID        uint      `json:"id"`
	Site      Site      `json:"site"`
	OpenLogin OpenLogin `json:"open_login"`
}

func (v Site) Value() (driver.Value, error)       { return json.Marshal(v) }
func (v *Site) Scan(value interface{}) error      { return json.Unmarshal(value.([]byte), v) }
func (v OpenLogin) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *OpenLogin) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }
