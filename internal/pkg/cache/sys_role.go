package cache

import (
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"time"
)

func GetSysRoleList(key string) []model.SysRole {
	redisdb := initialize.GetRedisIns()
	list := make([]model.SysRole, 0)
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
	var role model.SysRole
	for _, value := range values {
		util.Json2Struct(value, &role)
		list = append(list, role)
	}
	return list
}

/*
	Redis Lpush 命令将一个或多个值插入到列表头部，导致最后插入的在列表最前面
	Redis Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)
*/
func SetSysRoleList(key string, values []model.SysRole) error {
	redisdb := initialize.GetRedisIns()
	strs := make([]string, 0)
	for _, value := range values {
		strs = append(strs, util.Struct2Json(value))
	}
	// 设置key的过期时间
	redisdb.Expire(key, 2*time.Hour)
	return redisdb.RPush(key, strs).Err()

}
