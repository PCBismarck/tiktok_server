package redis_config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	Redisdb *redis.Client
	ctx     = context.Background()
)

func InitRedis() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "101.43.172.154:6379",
		Password: "douyinDY_123",
		DB:       0,
	})
}
