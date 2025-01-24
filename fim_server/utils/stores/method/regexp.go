package method

import (
	"regexp"
	"strings"
)

type regexpServerInterface interface {
	GetIP() []string // 匹配IP:端口
}
type regexpServer struct{ String string }

func Regexp(s string) regexpServerInterface { return &regexpServer{String: s} }

func (s *regexpServer) GetIP() []string {
	re := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}(:\d+)?`)
	matches := re.FindAllString(s.String, -1)
	
	var result []string
	for _, match := range matches {
		result = append(result, strings.TrimSpace(match)) // 清理掉多余的空格并返回结果
	}
	return result
}
