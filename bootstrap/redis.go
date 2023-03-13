package bootstrap

import (
	"bingo-example/pkg/cache"
	"bingo-example/pkg/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// SetupRedis 设置redis
func SetupRedis() {
	cache.Connect(&redis.Options{
		Addr: fmt.Sprintf("%s:%v",
			config.Get("redis.host"),
			config.Get("redis.port")),
		Password: config.Get("redis.password"),
		// 连接池容量及闲置连接数量
		PoolSize:     15,
		MinIdleConns: 10,
		// 命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: time.Millisecond * 8,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: time.Millisecond * 512, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
		// 超时
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		PoolTimeout:  time.Second * 3,
	})
}
