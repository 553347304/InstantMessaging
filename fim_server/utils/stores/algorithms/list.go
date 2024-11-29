package algorithms

import (
	"regexp"
)

func InList(arr []string, key string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
}

func InListRegex(arr []string, key string) bool {
	for _, s := range arr {
		regex, _ := regexp.Compile(s)
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}
