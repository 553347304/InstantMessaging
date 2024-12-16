package method

import (
	"fim_server/utils/stores/logs"
	"regexp"
)

// InList Key 是否在列表中
func InList(arr []string, key string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
}

// InListRegex Key 是否在列表中 支持正则表达式
func InListRegex(arr []string, key string) bool {
	for _, s := range arr {
		logs.Info(s)
		regex, err := regexp.Compile(s)
		if err != nil {
			logs.Info("正则表达式编译失败:", err)
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}

// Deduplication 去重
func Deduplication[T comparable](slice []T) []T {
	var result []T

	maps := make(map[T]bool)

	for _, v := range slice {
		if !maps[v] {
			maps[v] = true
			result = append(result, v)
		}
	}

	return result
}


// ListIntersect 求切片交集
func ListIntersect[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, v := range s1 { m[v] = true }

	var slice []T
	for _, v := range s2 {
		if m[v] {
			slice = append(slice, v)
		}
	}
	return slice
}

// ListDifference 求切片差集
func ListDifference[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, v := range s2 { m[v] = true }

	var slice []T
	for _, v := range s1 {
		if !m[v] {
			slice = append(slice, v)
		}
	}
	return slice
}
