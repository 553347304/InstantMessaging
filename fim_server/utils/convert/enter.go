package convert

import (
	"log"
	"strconv"
)

func Int(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Println("转换错误: " + err.Error())
		return -1
	}
	return number
}
