package main

import (
	"fim_server/config/core"
	"fim_server/config/flags"
	"fim_server/utils/stores/logs"
)

func main() {
	core.Init()
	flags.Command()
	logs.Info("初始化成功")

}
