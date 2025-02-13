package method

import (
	"fim_server/utils/stores/logs"
	"regexp"
	"strings"
)

type regexpServerInterface interface {
	GetIP() []string // 匹配IP:端口
	IsTel() bool     // 手机号
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

func (s *regexpServer) IsTel() bool {
	regex, err := regexp.Compile(`^1[3-9]\d{9}$`)
	if err != nil {
		logs.Warn("Error compiling regex:", err)
		return false
	}
	return regex.MatchString(s.String)
}
