package method_list

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"regexp"
)

func In(arr []string, key string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
} // Key是否在列表中
func InRegex(arr []string, key string) bool {
	for _, s := range arr {
		regex, err := regexp.Compile(s)
		if err != nil {
			logs.Info("正则表达式编译失败:", err)
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
} // Key是否在列表中   支持正则  包含
func Unique[T comparable](s1 []T) (slice []T) {
	m := make(map[T]bool)

	for _, v := range s1 {
		if !m[v] {
			m[v] = true
			slice = append(slice, v)
		}
	}

	return
} // 切片去重
func Intersect[T comparable](s1, s2 []T) (slice []T) {
	m := conv.SliceMap(s1)
	for _, v := range s2 {
		if m[v] {
			slice = append(slice, v)
		}
	}
	return
} // 切片交集
func Difference[T comparable](s1, s2 []T) (slice []T) {
	m := conv.SliceMap(s2)
	for _, v := range s1 {
		if !m[v] {
			slice = append(slice, v)
		}
	}
	return
} // 切片差集
