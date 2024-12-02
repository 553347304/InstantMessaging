package core

import (
	"fim_server/utils/src"
)

const (
	SystemLog  = "release" // 运行模式支持：gin.ReleaseMode|"release"|"test"|"debug"
	configFile = "./settings.yaml"
)

func Init() {
	src.Config = Yaml()
	src.DB = Mysql(src.Config.System.Mysql)
	src.Redis = Redis(src.Config.System.Redis)
}
