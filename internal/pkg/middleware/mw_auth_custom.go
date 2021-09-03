package middleware

import (
	"strings"

	"github.com/vagaryer/go-web/internal/pkg/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

/*
	基于jwt-go，自己写的简单的token验证中间件
*/

var jwtkey = []byte("go-web")

//token授权验证
func AuthMiddleware(redisIns *redis.Client, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Abort()
			return
		}
		tokenString = strings.SplitN(tokenString, " ", 2)[1]

		token, claims, err := parseToken(tokenString, redisIns)
		if err != nil || token == nil || !token.Valid {
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}

}

func parseToken(tokenString string, redisIns *redis.Client) (*jwt.Token, *model.Claims, error) {
	//判断token的有效性
	_, err := redisIns.Get(tokenString).Result()
	if err != nil {
		return nil, nil, err
	}
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})

	return token, claims, err
}
