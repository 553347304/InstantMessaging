package log_service

import (
	"gorm.io/gorm"
)

const (
	Action = "操作日志"
)

var DB *gorm.DB
