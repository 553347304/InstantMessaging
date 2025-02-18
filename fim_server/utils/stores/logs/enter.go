package logs

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type logServer struct{}

func (l *logServer) line() []string {
	// 获取当前根目录名
	cmd, err := exec.Command("go", "list", "-m").Output()
	if err != nil {
		log.Fatal("go list 获取失败", err)
	}
	dir := regexp.MustCompile(`\s+`).ReplaceAllString(string(cmd), "") + "/"
	
	// 获取路径行号
	var pathList = make([]string, 0)
	pc := make([]uintptr, 32)
	n := runtime.Callers(3, pc)
	cf := runtime.CallersFrames(pc[:n])
	
	for {
		frame, more := cf.Next()
		if !more {
			break
		} else {
			pathLine := fmt.Sprintf("%s:%d", frame.File, frame.Line)
			// 只返回当前项目的路径
			index := strings.Index(pathLine, dir)
			
			if index != -1 {
				pathLine = strings.Replace(pathLine[index:], dir, "", 1)
				pathList = append(pathList, pathLine)
			}
			
		}
	}
	
	return pathList
}
func (l *logServer) hexToRgb(hex string) (int, int, int) {
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
func (l *logServer) color(s any, hex ...string) string {
	s = fmt.Sprintf("%+v",s)
	if len(hex) >= 3 {
		if hex[2] == "斜体" {
			s = fmt.Sprintf("\033[3m%s", s) // 斜体
		}
	}
	r, g, b := l.hexToRgb(hex[0])
	if len(hex) >= 2 {
		br, bg, bb := l.hexToRgb(hex[1])
		s = fmt.Sprintf("\033[48;2;%d;%d;%dm%s", br, bg, bb, s) // 背景色
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, s)
}
func (l *logServer) time() string {
	return l.color(fmt.Sprintf("[%v]", time.Now().Format(c.Time)), "#ffffff")
}
func (l *logServer) string(s ...interface{}) string {
	source := ""
	for i := 0; i < len(s); i++ {
		source += " " + fmt.Sprint(s[i])
	}
	source = source[1:]
	return source
}
func (l *logServer) ip(source string) string {
	// 是否为IP地址
	regex := regexp.MustCompile(`.(\d{1,3}\.){3}\d{1,3}(:\d+)?`).FindAllString(source, -1)
	for _, ip := range regex {
		if strings.Index(ip, "/") == -1 {
			isIp := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}(:\d+)?`).FindString(ip)
			if isIp != "" {
				source = strings.ReplaceAll(source, isIp, l.color(isIp, "#FFFFFF", "#288F6A", "斜体"))
			}
		}
	}
	
	return source
}

func (l *logServer) error(line []string, s ...interface{}) {
	fmt.Println(l.time(), l.color(l.string(s...), c.Error))
	
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
}

var l = &logServer{}
