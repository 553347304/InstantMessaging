package main

import (
	"fim_server/config"
	"fim_server/config/flags"
	"fim_server/utils/stores/logs"
)

func main() {
	config.Init()
	
	flags.Command()
	logs.Info("初始化成功")
}
