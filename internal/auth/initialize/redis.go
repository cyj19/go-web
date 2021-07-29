package initialize

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	redisIns  *redis.Client
	onceRedis sync.Once
)

// 初始化Redis
func Redis() {

	// 单例模式，保证生命周期只初始化一次
	onceRedis.Do(func() {
		redisIns = redis.NewClient(&redis.Options{
			Addr:     box.ViperIns.GetString("redis.addr"),
			Password: box.ViperIns.GetString("redis.password"),
			DB:       box.ViperIns.GetInt("redis.db"),
		})
	})
	_, err := redisIns.Ping().Result()
	if err != nil || redisIns == nil {
		panic(fmt.Sprintf("初始化Redis异常：%v", err))
	}
}

// 暴露给其他包
func GetRedisIns() *redis.Client {
	return redisIns
}
