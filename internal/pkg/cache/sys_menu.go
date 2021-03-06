package cache

import (
	"time"

	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/util"

	"github.com/go-redis/redis"
)

func GetSysMenuList(redisdb *redis.Client, key string) []model.SysMenu {
	list := make([]model.SysMenu, 0)
	// 获取List长度
	rLen, err := redisdb.LLen(key).Result()
	if err != nil {
		return nil
	}
	values, err := redisdb.LRange(key, 0, rLen-1).Result()
	if err != nil {
		return nil
	}
	var menu model.SysMenu
	for _, value := range values {
		util.Json2Struct(value, &menu)
		list = append(list, menu)
	}

	return list
}

/*
	Redis Lpush 命令将一个或多个值插入到列表头部，导致最后插入的在列表最前面
	Redis Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)
*/
func SetSysMenuList(redisdb *redis.Client, key string, values []model.SysMenu) error {
	strs := make([]string, 0)
	for _, value := range values {
		strs = append(strs, util.Struct2Json(value))
	}
	// 设置过期时间
	redisdb.Expire(key, 2*time.Hour)
	return redisdb.RPush(key, strs).Err()
}
