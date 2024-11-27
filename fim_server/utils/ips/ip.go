package ips

import (
	"regexp"
	"strings"
)

// netx.InternalIp()  获取内网IP

// IsIP 提取字符串中的所有IP地址和端口
func IsIP(input string) []string {
	// 定义正则表达式：匹配IP地址和可选的端口
	// 匹配 IPv4 地址：x.x.x.x 格式
	// 可选的端口号格式：:端口号
	re := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}(:\d+)?`)
	matches := re.FindAllString(input, -1)

	var result []string
	for _, match := range matches {
		result = append(result, strings.TrimSpace(match)) // 清理掉多余的空格并返回结果
	}
	return result
}
