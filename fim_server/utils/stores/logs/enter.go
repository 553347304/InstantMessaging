package logs

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func hexToRgb(hex string) (int, int, int) {
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

func logColor(s any, hex ...string) string {
	s = fmt.Sprint(s)
	if len(hex) >= 3 {
		if hex[2] == "斜体" {
			s = fmt.Sprintf("\033[3m%s", s) // 斜体
		}
	}

	r, g, b := hexToRgb(hex[0])
	if len(hex) >= 2 {
		br, bg, bb := hexToRgb(hex[1])
		s = fmt.Sprintf("\033[48;2;%d;%d;%dm%s", br, bg, bb, s) // 背景色
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, s)
}

// getLine 获取路径行号
func getLine(skip int) string {
	_, callPath, callLine, _ := runtime.Caller(skip) // 获取调用者信息
	// 获取当前根目录名
	output, _ := exec.Command("go", "list", "-m").CombinedOutput()
	name := strings.TrimSpace(string(output)) + "/"
	index := strings.Index(callPath, name)
	path := strings.Replace(callPath[index:], name, "", 1)
	return fmt.Sprintf(" %s:%d ", path, callLine)
}

func isSprintf(s ...interface{}) string {

	str := fmt.Sprint(s[0])

	is := strings.Contains(str, "%s") || strings.Contains(str, "%d") || strings.Contains(str, "%v")

	message := ""
	for i, i2 := range s {
		message += fmt.Sprint(i2)
		if i < len(s)-1 {
			message += " "
		}
	}

	if is {
		message = fmt.Sprintf(str, s[1:]...)
	}
	return message
}

func isFieldColor(s string) string {
	var message = s

	// 是否为IP地址
	isIP := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}(:\d+)?`).FindAllString(s, -1)
	for _, match := range isIP {
		ip := strings.TrimSpace(match)
		color := logColor(ip, "#FFFFFF", "#288F6A", "斜体")
		message = strings.ReplaceAll(message, ip, color)
	}

	return message
}

const line = 2
const times = "15:04:05"
