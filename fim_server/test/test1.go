package main

import (
	"fim_server/utils/stores/method"
)

func main() {

	arr := []string{"created_at", "banana", "cherry"}
	a2 := method.List(arr).InRegex("created_at desc")
	method.
		logs.Info(a2)
}
