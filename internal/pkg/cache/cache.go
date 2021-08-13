package cache

import (
	"encoding/json"
	"fmt"
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/util"
	"time"

	"github.com/go-redis/redis"
)

/*
	使用Redis做缓存，定义key格式为 表名:字段名:字段值
	例如：user:id:1:username:vagaryer
*/
// model必须是指针类型
func Get(key string, model interface{}) error {
	redisdb := initialize.GetRedisIns()
	value, err := redisdb.Get(key).Result()
	if err == redis.Nil || err != nil {
		global.LoggerIns.Println("redis missing key err:", err)
		return err
	}
	// 缓存中查询
	fmt.Println("缓存中查询")
	// json 转 model
	return util.Json2Struct(value, model)

}

func Set(key string, value interface{}) error {
	redisdb := initialize.GetRedisIns()
	valueData, err := json.Marshal(value)
	if err != nil {
		global.LoggerIns.Println("json marshal err:", err)
		return err
	}

	return redisdb.Set(key, string(valueData), 2*time.Hour).Err()

}

func Exist(key string) bool {
	redisdb := initialize.GetRedisIns()
	num, err := redisdb.Exists(key).Result()
	if err != nil || num < 1 {
		return false
	}
	return true
}

// 删除key
func Del(key ...string) error {
	redisdb := initialize.GetRedisIns()
	return redisdb.Del(key...).Err()
}

// 返回与pattern匹配的key
func Keys(pattern string) []string {
	redisdb := initialize.GetRedisIns()
	return redisdb.Keys(pattern).Val()
}
