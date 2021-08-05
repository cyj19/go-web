package router

import (
	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/store"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// 注册用户路由
func InitUserRouter(r *gin.RouterGroup, factoryIns store.Factory, enforcer *casbin.Enforcer, authMiddleware *jwt.GinJWTMiddleware) {

	userv1 := r.Group("/user")
	userv1.Use(authMiddleware.MiddlewareFunc())
	//userv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(initialize.GetEnforcerIns()))
	{
		userHandler := user.NewSysUserHandler(factoryIns, enforcer)

		userv1.POST("/add", userHandler.Create)
		userv1.DELETE("/delete", userHandler.BatchDelete)
		userv1.PATCH("/update", userHandler.Update)
		userv1.PATCH("/role/update", userHandler.UpdateRoleForUser)
		userv1.POST("/page", userHandler.GetPage)

	}

}
