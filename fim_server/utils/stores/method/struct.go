package method

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"reflect"
)

type structServerInterface interface {
	Is(interface{}) bool                         // 是不是结构体
	To(interface{}, interface{}) bool            // 结构体转换类型   source | scan
	Delete(interface{}, ...string) []interface{} // 删除指定字段	结构体指针 | 字段   返回:删掉字段的值
}
type structServer struct {
	String string
}

func Struct() structServerInterface { return &structServer{} }

func (s *structServer) delete(scan interface{}, field string) interface{} {
	elem := reflect.ValueOf(scan)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
		logs.Info(elem)
	}
	if elem.Kind() != reflect.Struct {
		return nil
	}
	f := elem.FieldByName(field)
	if f.IsValid() && f.CanSet() {
		value := f.Interface()        // 获取字段值
		f.Set(reflect.Zero(f.Type())) // 删除指定字段
		return value
	}
	return nil
}
func (s *structServer) Is(v interface{}) bool {
	elem := reflect.ValueOf(v)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	return elem.Kind() == reflect.Struct
}
func (s *structServer) To(source interface{}, scan interface{}) bool {
	return conv.Json().Unmarshal(conv.Json().Marshal(source), &scan)
}
func (s *structServer) Delete(v interface{}, field ...string) (value []interface{}) {
	for i := 0; i < len(field); i++ {
		value = append(value, s.delete(v, field[i]))
	}
	return
}

// func StructMap(data interface{}) map[string]interface{} {
// 	maps := make(map[string]interface{})
// 	values := reflect.ValueOf(data)
// 	types := values.Type() // = reflect.TypeOf(data)
// 	if values.Kind() == reflect.Struct {
// 		for i := 0; i < values.NumField(); i++ {
// 			name := types.Field(i).Name     // 字段名
// 			t := types.Field(i).Type.Kind() // 字段类型
// 			v := values.Field(i)            // 字段值
// 			tag := types.Field(i).Tag       // 字段Tag
// 			logs.Info(name, t, v, tag)
// 		}
// 	}
// 	return maps
// }
