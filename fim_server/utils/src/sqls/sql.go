package sqls

func GetList[T any](model T, m Mysql) MysqlResponse[[]T] {
	var total int64
	var scan = make([]T, 0)
	m.Param(model)
	m.isFieldExist(model, m.PageInfo.Sort) // 排序字段是否存在
	m.DB.Model(model).Count(&total).Order(m.PageInfo.Sort).
		Limit(m.PageInfo.Limit).Offset(m.PageInfo.Page).Find(&scan)
	return MysqlResponse[[]T]{Total: total, List: scan}
}

func GetListGroup[T any, R any](model T, m Mysql, scan *[]R) MysqlResponse[[]R] {
	var total int64
	m.Param(model)
	m.isFieldExist(model, m.PageInfo.Sort) // 排序字段是否存在
	m.DB.Model(model).Count(&total).Order(m.PageInfo.Sort).
		Limit(m.PageInfo.Limit).Offset(m.PageInfo.Page).Find(&scan)
	return MysqlResponse[[]R]{Total: total, List: *scan}
}
