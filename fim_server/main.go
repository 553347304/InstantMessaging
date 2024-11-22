package main

import (
	"fim_server/global/core"
	"fim_server/global/flags"
	"fmt"
)

func main() {
	core.Init()
	flags.Command()
	fmt.Println("初始化成功")
}
