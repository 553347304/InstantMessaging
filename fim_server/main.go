package main

import (
	"fim_server/core"
	"fim_server/flags"
	"fmt"
)

func main() {
	core.Init()
	flags.Command()
	fmt.Println("初始化成功")
}
