package middleware

import (
	"strings"
	"sync"

	"github.com/cyj19/go-web/internal/admin/api/v1/user"
	"github.com/cyj19/go-web/internal/admin/store"
	"github.com/cyj19/go-web/internal/pkg/config"
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"
	"github.com/cyj19/go-web/internal/pkg/util"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var checkLock sync.Mutex

// CasbinMiddleware 基于rbac
func CasbinMiddleware(factory store.Factory, conf *config.Configuration, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

		obj := c.Request.URL.Path
		// 清除路径前缀
		obj = strings.Replace(obj, "/"+conf.Server.UrlPrefix, "", 1)
		act := c.Request.Method
		// 获取当前用户
		userHandler := user.NewSysUserHandler(factory)
		currentUser := userHandler.GetCurrentUser(c)

		if !check(enforcer, currentUser, obj, act) {
			c.Abort()
			response.FailWithCode(response.Forbidden)
			return
		}

		c.Next()
	}

}

func check(enforcer *casbin.Enforcer, user model.SysUser, obj string, act string) bool {
	checkLock.Lock()
	defer checkLock.Unlock()
	if len(user.Roles) <= 0 {
		return false
	}
	var flag int

	for i, role := range user.Roles {
		b, _ := enforcer.Enforce(util.Uint642Str(role.Id), obj, act)
		if b {
			return true
		}
		flag = i + 1
	}
	return flag != len(user.Roles)

}
