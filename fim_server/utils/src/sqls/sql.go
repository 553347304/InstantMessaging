package sqls

import "strings"

func GetList[T any](model T, m Mysql) ([]T, int64) {
	var total int64
	var list = make([]T, 0)
	m.Param(model)
	m.isFieldExist(model, strings.Split(m.PageInfo.Sort, " ")[0]) // 排序字段是否存在
	m.DB.Model(model).
		Count(&total).
		Order(m.PageInfo.Sort).
		Limit(m.PageInfo.Limit).
		Offset(m.PageInfo.Page).
		Find(&list)
	return list, total
}

func GetListGroup[T any, R any](model T, scan *[]R, m Mysql) int64 {
	var total int64
	m.Param(model)
	m.DB.Model(model).
		Count(&total).
		Order(m.PageInfo.Sort).
		Limit(m.PageInfo.Limit).
		Offset(m.PageInfo.Page).
		Find(&scan)
	return total
}
