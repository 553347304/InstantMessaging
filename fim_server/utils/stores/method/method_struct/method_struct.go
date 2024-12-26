package method_struct

import (
	"fim_server/utils/stores/conv"
)

// ReplaceStruct 替换结构体
func ReplaceStruct[T any](source any) T {
	var m = new(T)
	conv.Unmarshal(conv.Marshal(source), &m)
	return *m
}
