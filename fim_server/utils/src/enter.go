package src

import (
	"fim_server/config"
	"gorm.io/gorm"
)

type PageInfo struct {
	Key   string `form:"key"`   // 关键字
	Like  string `form:"like"`  // 模糊匹配
	Page  int    `form:"page"`  // 当前页数
	Limit int    `form:"limit"` // 每页条数
	Sort  string `form:"sort"`  // 排序---从后往前:"_ desc"   从前往后:"_ asc"
}

func (o *PageInfo) Param() {
	// 默认条数
	if o.Limit == 0 {
		o.Limit = -1 // 全部
	}
	// 默认页数
	if o.Page != 0 {
		o.Page = (o.Page - 1) * o.Limit
	}
	// 默认排序
	// if o.Sort == "" {
	// 	o.Sort = "created_at desc"
	// }
}

var (
	Config *config.Config
	DB     *gorm.DB
	// Redis  *redis.Client
)
