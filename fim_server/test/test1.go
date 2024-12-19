package main

import (
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method/method_list"
)

func main() {

	arr := []string{"created_at", "banana", "cherry"}
	a2 := method_list.InRegex(arr, "created_at desc")



	logs.Info(a2)
}

