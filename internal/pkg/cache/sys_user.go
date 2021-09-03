package cache

import (
	"time"

	"github.com/vagaryer/go-web/internal/pkg/model"
	"github.com/vagaryer/go-web/internal/pkg/util"

	"github.com/go-redis/redis"
)

func GetSysUserList(redisdb *redis.Client, key string) []model.SysUser {
	list := make([]model.SysUser, 0)
	rLen, err := redisdb.LLen(key).Result()
	if err != nil {
		// 写入日志
		return nil
	}
	values, err := redisdb.LRange(key, 0, rLen-1).Result()
	if err != nil {
		// 写入日志
		return nil
	}
	var user model.SysUser
	for _, value := range values {
		util.Json2Struct(value, &user)
		list = append(list, user)
	}
	return list
}

/*
	Redis Lpush 命令将一个或多个值插入到列表头部，导致最后插入的在列表最前面
	Redis Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)
*/
func SetSysUserList(redisdb *redis.Client, key string, values []model.SysUser) error {
	strs := make([]string, 0)
	for _, value := range values {
		strs = append(strs, util.Struct2Json(value))
	}
	// 设置key的过期时间
	redisdb.Expire(key, 2*time.Hour)
	return redisdb.RPush(key, strs).Err()
}
