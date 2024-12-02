package sqls

import (
	"fim_server/utils/src"
	"fim_server/utils/stores/logs"
	"fmt"
	"strings"
)

func (m *Mysql) isFieldExist(model interface{}, fieldName string) bool {
	if fieldName == "" {
		return true
	}
	is := m.DB.Migrator().HasColumn(model, fieldName)
	if !is {
		tableName := fmt.Sprintf("%T", model)
		logs.Error(fmt.Sprintf("'%s' 表中不存在 '%s' 字段", tableName, fieldName))
	}
	return is
}

func (m *Mysql) Param(model interface{}) bool {
	// 默认数据库
	if m.DB == nil {
		m.DB = src.DB
	}

	is1 := m.isFieldExist(model, m.PageInfo.Search)
	is2 := m.isFieldExist(model, strings.Split(m.PageInfo.Sort, " ")[0])
	if !is1 || !is2 {
		return false
	}

	m.DB = m.DB.Where(model) // 查结构体自身条件

	// 预加载
	for _, preload := range m.Preload {
		m.DB = m.DB.Preload(preload)
	}

	return true
}

// // Where 匹配
// if m.PageInfo.Search != "" && m.PageInfo.Key != "" {
// 	switch m.PageInfo.Like {
// 	case "like":
// 		m.DB = m.DB.Where(m.PageInfo.Search+" like ?", "%"+m.PageInfo.Key+"%") // 模糊匹配
// 	default:
// 		m.DB = m.DB.Where(m.PageInfo.Search+" = ?", m.PageInfo.Key) // 精确匹配
// 	}
// }
