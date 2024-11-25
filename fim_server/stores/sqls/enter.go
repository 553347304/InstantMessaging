package sqls

import "fim_server/global"

func Take[T any](m *T, dest interface{}, arms ...interface{}) bool {
	return global.DB.Take(m, dest, arms).Error == nil
}

func Take[T any](m *T, dest interface{}, arms ...interface{}) bool {
	return global.DB.Where(m, dest, arms)
}

func Take[T any](m *T, q interface{}, args ...interface{}) bool {
	var total int64
	global.DB.Model(m).Where(q, args...).Count(&total).Find(&m)
	return total != 0
}
