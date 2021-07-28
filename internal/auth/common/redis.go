package common

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	redisIns  *redis.Client
	onceRedis sync.Once
)

//单例
func GetRedisIns() (*redis.Client, error) {
	var err error
	onceRedis.Do(func() {
		redisIns = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.addr"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		})
	})
	_, err = redisIns.Ping().Result()
	return redisIns, err
}
