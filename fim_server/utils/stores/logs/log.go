package logs

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Info(s ...interface{}) string {
	str := isSprintf(s...)
	t := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	fmt.Println(t + getLine(line) + logColor(isFieldColor(str), "#80ffff"))
	return str
}

func Error(s ...interface{}) error {
	str := isSprintf(s...)
	t := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	fmt.Println(t + getLine(line) + logColor(isFieldColor(str), "#ff0000"))

	return errors.New(str)
}

func Structs(s interface{}) {
	t := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	fmt.Println(t + getLine(line) + logColor("结构体", "#ff0000"))
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
			fmt.Printf(fmt.Sprintf("%s: %v", logColor(fieldName, "#ff0000"), logColor(field.Interface(), "#00ffff")))
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
