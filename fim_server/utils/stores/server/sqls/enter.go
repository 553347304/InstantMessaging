package sqls

import (
	"fim_server/utils/stores"
	"gorm.io/gorm"
)

type Mysql struct {
	DB       *gorm.DB
	Preload  []string
	PageInfo stores.PageInfo
}
