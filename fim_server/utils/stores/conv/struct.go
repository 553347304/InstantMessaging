package conv

import (
	"reflect"
)

type structServerInterface[T interface{}] interface {
	Type(interface{}) T                                // 同字段 替换结构体类型
	StructMap(...string) map[string]interface{}        // 结构体转map返回指定字段 | 不填:返回全部
	StructSliceMap(...string) []map[string]interface{} // 结构体数组转map返回指定字段 | 不填:返回全部
}
type structServer[T interface{}] struct{ Struct T }

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Struct[T interface{}](m T) structServerInterface[T] { return &structServer[T]{Struct: m} }

func (s *structServer[T]) Type(source interface{}) T {
	var m = new(T)
	Json().Unmarshal(Json().Marshal(source), &m)
	return *m
}
func (s *structServer[T]) StructMap(contain ...string) map[string]interface{} {
	var maps = make(map[string]interface{})
	var scan map[string]interface{}
	if !Json().Unmarshal(Json().Marshal(s.Struct), &scan) {
		return maps
	}
	// 不填返回全部字段
	if len(contain) == 0 {
		return scan
	}
	// 返回指定字段
	for key, value := range scan {
		for i := 0; i < len(contain); i++ {
			if contain[i] == key {
				maps[key] = value
			}
		}
	}
	return maps
}
func (s *structServer[T]) StructSliceMap(contain ...string) []map[string]interface{} {
	var maps = make([]map[string]interface{}, 0)
	value := reflect.ValueOf(s.Struct)
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			v := value.Index(i).Interface()
			_map := Struct(v).StructMap(contain...)
			maps = append(maps, _map)
		}
	}
	return maps
}

/*
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
				continue
			}

			value := v.Interface()
			if t == reflect.Ptr {
				value = v.Elem().Interface() // Elem 解指针引用
			}

			// 遍历子结构体
			returnMap := StructMap(value, _tag)
			if len(returnMap) != 0 {
				maps[tag] = returnMap
				continue
			}

			// 遍历子结构体切片
			returnStructSlice := structMapSlice(value, _tag)
			if len(returnStructSlice) != 0 {
				maps[tag] = returnStructSlice
				continue
			}

			maps[tag] = value
		}
	}
	return maps
}
func structMapSlice(data interface{}, _tag string) []map[string]interface{} {
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
*/
