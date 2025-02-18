package method

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"reflect"
	"regexp"
)

type structServerInterface interface {
	GetValue(interface{}) reflect.Value
	Delete(interface{}, ...string) []interface{} // 删除指定字段	结构体指针 | 字段   返回:删掉字段的值
	To(interface{}, interface{}) bool            // 结构体转换类型   source | scan
	ToMap(interface{}) map[string]interface{}
	ToMapSlice(interface{}) []map[string]interface{}
}
type structServer struct {
	String string
}

func Struct() structServerInterface { return &structServer{} }

func (s *structServer) delete(v reflect.Value, field string) interface{} {
	f := v.FieldByName(field)
	if f.IsValid() && f.CanSet() {
		value := f.Interface()        // 获取字段值
		f.Set(reflect.Zero(f.Type())) // 删除指定字段
		return value
	}
	return nil
}
func (s *structServer) GetValue(v interface{}) reflect.Value {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
func (s *structServer) Delete(v interface{}, field ...string) (value []interface{}) {
	_value := s.GetValue(v)
	if _value.Kind() != reflect.Struct {
		logs.Error("不是结构体")
		return nil
	}
	
	for i := 0; i < len(field); i++ {
		value = append(value, s.delete(_value, field[i]))
	}
	
	return
}
func (s *structServer) To(source interface{}, scan interface{}) bool {
	return conv.Json().Unmarshal(conv.Json().Marshal(source), &scan)
}
func (s *structServer) ToMap(v interface{}) map[string]interface{} {
	scan := make(map[string]interface{})
	_value := s.GetValue(v)
	if _value.Kind() != reflect.Struct {
		logs.Error("不是结构体", v)
		return scan
	}
	_type := _value.Type()
	
	for i := 0; i < _value.NumField(); i++ {
		field := _value.Field(i)
		fieldType := _type.Field(i)
		
		// Tag
		jsonTag := fieldType.Tag.Get("json")
		tag := regexp.MustCompile(`-|,optional|\s+`).ReplaceAllString(jsonTag, "")
		if tag == "" {
			continue
		}
		
		// 忽略 零值 | 指针空值
		if field.IsZero() {
			continue
		}
		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		
		scan[tag] = field.Interface()
		
		if field.Kind() == reflect.Struct {
			scan[tag] = s.ToMap(field.Interface())
		}
		
		if field.Kind() == reflect.Slice {
			scan[tag] = s.ToMapSlice(field.Interface())
		}
	}
	return scan
}
func (s *structServer) ToMapSlice(v interface{}) []map[string]interface{} {
	elem := reflect.ValueOf(v)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	
	scan := make([]map[string]interface{}, 0)
	if elem.Kind() == reflect.Slice {
		for j := 0; j < elem.Len(); j++ {
			item := elem.Index(j)
			if item.Kind() == reflect.Struct {
				maps := s.ToMap(item.Interface())
				scan = append(scan, maps)
			}
		}
	}
	return scan
}
