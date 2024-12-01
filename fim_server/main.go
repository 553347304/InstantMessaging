package main

import (
	"fim_server/config/core"
	"fim_server/config/flags"
	"fmt"
)

func main() {
	core.Init()
	flags.Command()
	fmt.Println("初始化成功")

}
