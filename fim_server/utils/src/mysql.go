package src

import (
	"fim_server/utils/stores/logs"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// T = scan result type
type serviceInterfaceMysql[T interface{}] interface {
	GetList() serviceMysqlResponse[[]T]      // 获取列表
	GetListGroup() serviceMysqlResponse[[]T] // 获取分组查询列表
}
type ServiceMysql[T interface{}] struct {
	DB       *gorm.DB
	Model    interface{} // 不填 Model 实例T类型作为Model
	Preload  []string    // 预加载
	PageInfo PageInfo    // 分页信息
	Where    string      // 模糊匹配
}
type serviceMysqlResponse[R any] struct {
	Total int64
	List  R
}

func (m *ServiceMysql[T]) where() {
	if m.PageInfo.Key == "" {
		return
	}
	split := strings.Split(m.Where, " ?")
	var quest []interface{}
	for _, s := range split {
		if strings.ReplaceAll(s, " ", "") != "" {
			key := m.PageInfo.Key
			if strings.Contains(s, "like") {
				key = fmt.Sprintf("%%%s%%", m.PageInfo.Key)
			}
			quest = append(quest, key)
		}
	}
	m.DB = m.DB.Where(m.Where, quest...)
}
func (m *ServiceMysql[T]) preload() {
	for _, preload := range m.Preload {
		m.DB = m.DB.Preload(preload)
	}
}
func (m *ServiceMysql[T]) param() {
	m.PageInfo.param()
	m.DB = m.DB.Where(m.Model) // 查结构体自身条件
	m.where()                  // 模糊匹配
	m.preload()                // 预加载
}
func (m *ServiceMysql[T]) isFieldExist() bool {
	fieldName := strings.Split(m.PageInfo.Sort, " ")[0] // 只取第一个字符串
	if fieldName == "" {
		return true
	}
	is := m.DB.Migrator().HasColumn(m.Model, fieldName)
	if !is {
		tableName := fmt.Sprintf("%T", m.Model)
		logs.Error(fmt.Sprintf("'%s' 表中不存在 '%s' 字段", tableName, fieldName))
	}
	return is
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Mysql[T interface{}](m ServiceMysql[T]) serviceInterfaceMysql[T] {
	if m.DB == nil {
		logs.Fatal("数据库不存在")
	}
	if m.Model == nil {
		m.Model = new(T)
	}
	return &ServiceMysql[T]{
		DB:       m.DB,
		Model:    m.Model,
		Preload:  m.Preload,
		PageInfo: m.PageInfo,
		Where:    m.Where,
	}
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func (m *ServiceMysql[T]) GetList() serviceMysqlResponse[[]T] {
	var total int64
	var scan = make([]T, 0)
	m.param()
	m.isFieldExist() // 排序字段是否存在
	m.DB.Model(m.Model).Count(&total).Order(m.PageInfo.Sort).
		Limit(m.PageInfo.Limit).Offset(m.PageInfo.Page).Find(&scan)
	return serviceMysqlResponse[[]T]{Total: total, List: scan}
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func (m *ServiceMysql[T]) GetListGroup() serviceMysqlResponse[[]T] {
	var total int64
	var scan = make([]T, 0)
	m.param()
	m.DB.Model(m.Model).Count(&total).Order(m.PageInfo.Sort).
		Limit(m.PageInfo.Limit).Offset(m.PageInfo.Page).Find(&scan)
	return serviceMysqlResponse[[]T]{Total: total, List: scan}
}
