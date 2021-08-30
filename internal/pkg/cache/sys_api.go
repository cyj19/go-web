package cache

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"time"

	"github.com/go-redis/redis"
)

func GetSysApiList(redisdb *redis.Client, key string) []model.SysApi {

	list := make([]model.SysApi, 0)
	rLen, err := redisdb.LLen(key).Result()
	if err != nil {
		return nil
	}

	values, err := redisdb.LRange(key, 0, rLen-1).Result()
	if err != nil {
		return nil
	}
	var api model.SysApi
	for _, value := range values {
		util.Json2Struct(value, &api)
		list = append(list, api)
	}

	return list
}

/*
	Redis Lpush 命令将一个或多个值插入到列表头部，导致最后插入的在列表最前面
	Redis Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)
*/
func SetSysApiList(redisdb *redis.Client, key string, values []model.SysApi) error {
	strs := make([]string, 0)
	for _, value := range values {
		strs = append(strs, util.Struct2Json(value))
	}
	// 设置过期时间
	redisdb.Expire(key, 2*time.Hour)
	return redisdb.RPush(key, strs).Err()
}
