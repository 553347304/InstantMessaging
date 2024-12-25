package conv

import (
	"fim_server/utils/stores/logs"
	"fmt"
	"reflect"
	"strconv"
)


type sliceServerInterface interface {
	String() []string
	Int() []int
}
type sliceServer struct {
	Slice []string
}
//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Slice(slice interface{}) sliceServerInterface {
	// reflection to string slice
	kind := reflect.TypeOf(slice).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		value := reflect.ValueOf(slice)
		s := make([]string, value.Len())
		for i := 0; i < value.Len(); i++ {
			s[i] = fmt.Sprint(value.Index(i).Interface())
		}
		return &sliceServer{Slice: s}
	}
	logs.Info("切片类型错误", kind)
	return &sliceServer{Slice: make([]string, 0)}
}

func (s *sliceServer) String () []string {
	return s.Slice
}
func (s *sliceServer) Int () []int {
	slice := make([]int, len(s.Slice))
	for i, str := range s.Slice {
		intVal, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			logs.Info("转换错误: %s\n", err)
			continue
		}
		slice[i] = int(intVal)
	}
	return slice
}
