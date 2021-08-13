package middleware

import "github.com/gin-gonic/gin"

type SkipperFunc func(*gin.Context) bool

//检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPreFixSkipper(prefixs ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixs {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}
