package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册基础路由，登录、登出、刷新token
func InitBaseRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	basev1 := r.Group("/base")
	{
		basev1.POST("/login", authMiddleware.LoginHandler)
		basev1.GET("/logout", authMiddleware.LogoutHandler)
		basev1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
}
