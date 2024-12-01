package src

import (
	"fim_server/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
