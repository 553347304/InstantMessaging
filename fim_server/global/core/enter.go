package core

import (
	"fim_server/global"
)

const (
	SystemLog  = "release" // 运行模式支持：gin.ReleaseMode|"release"|"test"|"debug"
	configFile = "settings.yaml"
)

func Init() {
	global.Config = Yaml()
	global.DB = Mysql(global.Config.System.Mysql)
	global.Redis = Redis(global.Config.System.Redis)
}
