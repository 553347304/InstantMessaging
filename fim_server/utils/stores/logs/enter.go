package logs

import (
	"fim_server/utils/stores/converts"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func logColor(hex string, s any) string {
	r, g, b := converts.HexRgb(hex)
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

const line = 2
const times = "15:04:05"
