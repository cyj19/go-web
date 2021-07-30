package middleware

import (
	"errors"
	"fmt"

	"go-web/internal/pkg/util"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		m := c.Request.Method
		userId := c.MustGet("user")

		b, err := enforcer.Enforce(fmt.Sprintf("%d", userId), p, m)
		if err != nil {
			util.WriteResponse(c, 500, err, nil)
			c.Abort()
			return
		} else if !b {
			util.WriteResponse(c, 403, errors.New("无权限访问"), nil)
			c.Abort()
			return
		}
		c.Next()
	}

}
