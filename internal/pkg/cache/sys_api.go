package cache

import (
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"time"
)

func GetSysApiList(key string) []model.SysApi {

	redisdb := initialize.GetRedisIns()
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

func SetSysApiList(key string, values []model.SysApi) error {
	redisdb := initialize.GetRedisIns()
	strs := make([]string, 0)
	for _, value := range values {
		strs = append(strs, util.Struct2Json(value))
	}
	// 设置过期时间
	redisdb.Expire(key, 2*time.Hour)
	return redisdb.LPush(key, strs).Err()
}
