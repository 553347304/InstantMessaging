package logs

import (
	"fmt"
	"reflect"
	"time"
)

func Info(s ...interface{}) {
	t := logColor("#ffffff", fmt.Sprintf("[%v]", time.Now().Format(times)))
	fmt.Println(t + getLine(line) + logColor("#80ffff", fmt.Sprint(s...)))
}

func Error(s ...interface{}) {
	t := logColor("#ffffff", fmt.Sprintf("[%v]", time.Now().Format(times)))
	fmt.Println(t + getLine(line) + logColor("#ff0000", fmt.Sprint(s...)))
}

func Structs(s interface{}) {
	t := logColor("#ffffff", fmt.Sprintf("[%v]", time.Now().Format(times)))
	fmt.Println(t + getLine(line) + logColor("#ff0000", "结构体"))
	value := reflect.ValueOf(s)
	subStructSlice(value)
	subStruct(value)
}
func subStructSlice(val reflect.Value) bool {
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			is1 := subStruct(val.Index(i))
			is2 := subStructSlice(val.Index(i))
			if is1 || is2 {
				continue
			}
			return true
		}
		return false
	}
	return false
}
func subStruct(val reflect.Value) bool {
	typ := val.Type()
	if val.Kind() == reflect.Struct {
		// 遍历结构体字段
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name
			fmt.Printf(fmt.Sprintf("%s: %v", logColor("#ff0000", fieldName), logColor("#00ffff", field.Interface())))
			is1 := subStruct(field)
			is2 := subStructSlice(field)
			if is1 || is2 {
				continue
			}
		}
		return true
	}
	return false
}
