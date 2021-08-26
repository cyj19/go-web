package middleware

import (
	"strings"

	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/model"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// 基于rbac
func CasbinMiddleware(factory store.Factory, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

		obj := c.Request.URL.Path
		// 清除路径前缀
		obj = strings.Replace(obj, "/"+global.Conf.Server.UrlPrefix, "", 1)
		act := c.Request.Method
		// 获取当前用户
		userHandler := user.NewSysUserHandler(factory, enforcer)
		currentUser := userHandler.GetCurrentUser(c)

		if !check(enforcer, currentUser, obj, act) {
			c.Abort()
			return
		}

		c.Next()
	}

}

func check(enforcer *casbin.Enforcer, user model.SysUser, obj string, act string) bool {
	if len(user.Roles) <= 0 {
		return false
	}
	var flag int

	for i, role := range user.Roles {
		b, _ := enforcer.Enforce(role.Id, obj, act)
		if b {
			return true
		}
		flag = i
	}
	if flag == len(user.Roles) {
		return false
	} else {
		return true
	}

}
