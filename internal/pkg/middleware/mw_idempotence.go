package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/vagaryer/go-web/internal/pkg/response"
)

/*
	幂等性中间件--防止重复提交
*/

var (
	expire = 24 * time.Hour
)

// redis lua脚本，必须以删除成功与否作为标志
const lua string = `
local current = redis.call('GET', KEYS[1])
if current == false then
	return '0';
end
local del = redis.call('DEl', KEYS[1])
if del == 1 then
	return '1';
else 
	return '0';
end
`

// 全局幂等性中间件
func Idempotence(redisIns *redis.Client, key string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 优先从请求头获取
		token := c.Request.Header.Get(key)
		if token == "" {
			// 从cookie中获取
			token, _ = c.Cookie(key)
		}
		token = strings.TrimSpace(token)
		if token == "" {
			c.Abort()
			response.FailWithMsg(response.IdempotenceTokenEmptyMsg)
		}
		// 校验token
		if !checkIdempotenceToken(redisIns, token) {
			c.Abort()
			response.FailWithMsg(response.IdempotenceTokenInvalidMsg)
		}
		c.Next()
	}
}

func checkIdempotenceToken(redisIns *redis.Client, token string) bool {
	// 执行脚本，删除token
	result, err := redisIns.Eval(lua, []string{token}).String()
	if err != nil || result != "1" {
		return false
	}
	return true
}

// 幂等性token获取接口
func GetIdempotenceToken(redisIns *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.SuccessWithData(GenIdempotenceToken(redisIns))
	}
}

func GenIdempotenceToken(redisIns *redis.Client) string {
	token := uuid.NewV4().String()
	// 写入redis
	redisIns.Set(token, true, expire)
	return token
}
