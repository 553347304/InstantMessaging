package main

import (
	"fim_server/config"
	"fim_server/config/flags"
	"fim_server/utils/stores/logs"
)

type UserInfo struct {
	Username string `json:"username"`
	Sign     string `json:"sign"`
	Avatar   string `json:"avatar"`
}
type T struct {
	UserInfo []UserInfo `json:"user_info"`
}

func main() {
	config.Init()
	
	flags.Command()
	logs.Info("初始化成功")
}
