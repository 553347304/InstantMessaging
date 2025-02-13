package method

import (
	"fim_server/utils/stores/logs"
	"reflect"
)

func StructMap(data interface{}) map[string]interface{} {
	maps := make(map[string]interface{})
	values := reflect.ValueOf(data)
	types := values.Type() // = reflect.TypeOf(data)
	if values.Kind() == reflect.Struct {
		for i := 0; i < values.NumField(); i++ {
			name := types.Field(i).Name     // 字段名
			t := types.Field(i).Type.Kind() // 字段类型
			v := values.Field(i)            // 字段值
			tag := types.Field(i).Tag       // 字段Tag
			logs.Info(name, t, v, tag)
		}
	}
	return maps
}
