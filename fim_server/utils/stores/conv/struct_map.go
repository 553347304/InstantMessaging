package conv

import (
	"encoding/json"
	"fim_server/utils/stores/logs"
	"reflect"
	"strings"
)

func StructJsonMap[T map[string]interface{} | []map[string]interface{}](m interface{}) T {
	var maps T
	jsonData, _ := json.Marshal(m)
	err := json.Unmarshal(jsonData, &maps)
	if err != nil {
		logs.Error(err)
	}
	return maps
}

// StructMap 结构体转Map
func StructMap(data interface{}, _tag string) map[string]interface{} {
	maps := make(map[string]interface{})
	values := reflect.ValueOf(data)
	types := values.Type() // = reflect.TypeOf(data)
	if values.Kind() == reflect.Struct {
		for i := 0; i < values.NumField(); i++ {
			name := types.Field(i).Name         // 字段名
			t := types.Field(i).Type.Kind()     // 字段类型
			v := values.Field(i)                // 字段值
			tag := types.Field(i).Tag.Get(_tag) // 字段Tag

			// 只提取一个json字段
			if strings.Contains(tag, ",") {
				tag = strings.Split(tag, ",")[0]
			}

			// 空值返回结构体字段
			if _tag == "" {
				tag = name
			}

			if tag == "*" {
				tag = types.Field(i).Tag.Get("json")
			}

			// 空值
			if v.IsZero() || tag == "" || tag == "-" {
				if isLog {
					logs.Error(tag, v)
				}
				continue
			}

			value := v.Interface()
			if t == reflect.Ptr {
				value = v.Elem().Interface() // Elem 解指针引用
			}

			if isLog {
				logs.Info(tag, value)
			}

			// 遍历子结构体
			returnMap := StructMap(value, _tag)
			if len(returnMap) != 0 {
				maps[tag] = returnMap
				continue
			}

			// 遍历子结构体切片
			returnStructSlice := StructMapSlice(value, _tag)
			if len(returnStructSlice) != 0 {
				maps[tag] = returnStructSlice
				continue
			}

			maps[tag] = value
		}
	}
	return maps
}

func StructMapSlice(data interface{}, _tag string) []map[string]interface{} {
	maps := make([]map[string]interface{}, 0)
	values := reflect.ValueOf(data)
	if values.Kind() == reflect.Slice {
		for i := 0; i < values.Len(); i++ {
			returnMap := StructMap(values.Index(i), _tag)
			if len(returnMap) != 0 {
				maps = append(maps, returnMap)
				continue
			}
		}
	}
	return maps
}
