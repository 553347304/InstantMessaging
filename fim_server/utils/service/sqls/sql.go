package sqls

func GetList[T any](model T, m Mysql) ([]T, int64) {
	var total int64
	var list = make([]T, 0)
	m.PageInfo.Param()
	if !m.Param(model) {
		return list, total
	}
	m.DB.Model(model).Count(&total).Limit(m.PageInfo.Limit).Offset(m.PageInfo.Page).Order(m.PageInfo.Sort).Find(&list)
	return list, total
}
