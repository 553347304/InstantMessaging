package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func Mysql(config string) *gorm.DB {
	var mysqlLogger = logger.Default.LogMode(logger.Error) // 日志等级
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalln("MySQL连接失败", config)
		return db
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout

	return db
}
