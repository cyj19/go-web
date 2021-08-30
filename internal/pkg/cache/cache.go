package cache

import (
	"encoding/json"
	"go-web/internal/pkg/util"
	"time"

	"github.com/go-redis/redis"
)

/*
	使用Redis做缓存，定义key格式为 表名:字段名:字段值
	例如：user:id:1:username:vagaryer
*/
// model必须是指针类型
func Get(redisdb *redis.Client, key string, model interface{}) error {
	value, err := redisdb.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}

	// json 转 model
	return util.Json2Struct(value, model)

}

func Set(redisdb *redis.Client, key string, value interface{}) error {
	valueData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return redisdb.Set(key, string(valueData), 2*time.Hour).Err()

}

func Exist(redisdb *redis.Client, key string) bool {
	num, err := redisdb.Exists(key).Result()
	if err != nil || num < 1 {
		return false
	}
	return true
}

// 删除key
func Del(redisdb *redis.Client, key ...string) error {
	return redisdb.Del(key...).Err()
}

// 返回与pattern匹配的key
func Keys(redisdb *redis.Client, pattern string) []string {
	return redisdb.Keys(pattern).Val()
}
