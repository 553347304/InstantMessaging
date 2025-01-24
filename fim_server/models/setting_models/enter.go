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
type ConfigModel struct {
	ID   uint `json:"id"`
	Site Site `json:"site"`
}

func (v Site) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *Site) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }
