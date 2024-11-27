package logs

import (
	"fim_server/utils/stores/converts"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func logColor(s any, hex ...string) string {
	if len(hex) >= 3 {
		if hex[2] == "斜体" {
			s = fmt.Sprintf("\033[3m%s", s) // 斜体
		}
	}

	r, g, b := converts.HexRgb(hex[0])
	if len(hex) >= 2 {
		br, bg, bb := converts.HexRgb(hex[1])
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

	message := fmt.Sprint(s...)
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
