package global

import (
	"fim_server/global/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
