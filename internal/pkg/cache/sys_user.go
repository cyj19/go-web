package cache

import (
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"time"
)

func GetSysUserList(key string) []model.SysUser {
	redisdb := initialize.GetRedisIns()
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

func SetSysUserList(key string, values []model.SysUser) error {
	redisdb := initialize.GetRedisIns()
	strs := make([]string, 0)
	for _, value := range values {
		strs = append(strs, util.Struct2Json(value))
	}
	// 设置key的过期时间
	redisdb.Expire(key, 2*time.Hour)
	return redisdb.LPush(key, strs).Err()
}
