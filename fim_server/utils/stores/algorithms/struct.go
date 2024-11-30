package algorithms

import "reflect"

func areStructsEqual(a, b interface{}) bool {
	// 使用反射获取结构体的值和类型
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// 如果不是结构体，返回 false
	if valA.Kind() != reflect.Struct || valB.Kind() != reflect.Struct {
		return false
	}

	// 比较两个结构体的字段数量
	if valA.NumField() != valB.NumField() {
		return false
	}

	// 遍历字段并比较每个字段的值
	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.Field(i)

		// 如果是指针字段，需要解引用比较
		if fieldA.Kind() == reflect.Ptr && fieldB.Kind() == reflect.Ptr {
			if fieldA.IsNil() && fieldB.IsNil() {
				continue // 都是 nil，跳过
			} else if fieldA.IsNil() || fieldB.IsNil() {
				return false // 其中一个是 nil，另一个不是
			}

			// 解引用后比较实际值
			if !reflect.DeepEqual(fieldA.Elem().Interface(), fieldB.Elem().Interface()) {
				return false
			}
		} else {
			// 对于非指针字段，直接比较
			if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
				return false
			}
		}
	}

	return true
}
