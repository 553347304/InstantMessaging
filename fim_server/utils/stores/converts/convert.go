package converts

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

func HexRgb(hex string) (int, int, int) {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return 0, 0, 0
	}
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return int(r), int(g), int(b)
}
