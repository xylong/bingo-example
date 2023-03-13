package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// Redis 获取redis
func Redis() *redis.Client {
	return rdb
}

// Connect 连接redis
func Connect(options *redis.Options) {
	rdb = redis.NewClient(options)

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
}
