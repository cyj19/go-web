package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/vagaryer/go-web/internal/admin/global"
	"github.com/vagaryer/go-web/internal/pkg/middleware"
)

// 注册基础路由，登录、登出、刷新token
func InitBaseRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	basev1 := r.Group("/base")
	{
		basev1.POST("/login", authMiddleware.LoginHandler)
		basev1.GET("/logout", authMiddleware.LogoutHandler)
		basev1.GET("/refresh_token", authMiddleware.RefreshHandler)
		// 幂等性token接口需要鉴权
		basev1.Use(authMiddleware.MiddlewareFunc()).GET("/idempotent_token", middleware.GetIdempotenceToken(global.RedisIns))
	}
}
