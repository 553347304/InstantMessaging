package core

import (
	"fim_server/utils/stores/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Mysql(config string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error), // 日志等级
	})
	if err != nil {
		logs.Fatal("MySQL连接失败", config)
		return db
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout

	return db
}
