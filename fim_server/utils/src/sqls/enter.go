package sqls

import (
	"fim_server/utils/src"
	"gorm.io/gorm"
)

type Mysql struct {
	DB       *gorm.DB
	Preload  []string     // 预加载
	PageInfo src.PageInfo // 分页信息
}

// Table 	指定表名   "user_models"
// Model 	指定表名   查询结构体
// Select 	返回字段	  gorm:"column:user"
// Joins 	连接类型
// Where 	高级查询
// Group 	分组
// DB.Raw(``).Find(&r)	// 原生sql
