package src

import (
	"github.com/go-redis/redis"
	"time"
)

type redisServerInterface interface {
	Set(string, interface{}, time.Duration)
	Del(string)
	IsExpiration(string) time.Duration // 判断key是否过期 | 过期:<0 | 单位/秒
}
type redisServer struct{ Redis *redis.Client }

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Redis(redis *redis.Client) redisServerInterface { return &redisServer{Redis: redis} }

func (s *redisServer) Set(key string, value interface{}, expiration time.Duration) {
	s.Redis.Set(key, value, expiration)
}
func (s *redisServer) Del(key string) {
	s.Redis.Del(key)
}

func (s *redisServer) IsExpiration(key string) time.Duration {
	t, err := s.Redis.TTL(key).Result()
	if err != nil {
		return 0
	}
	return t
}
