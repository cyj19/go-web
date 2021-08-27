package middleware

import (
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/response"
	"go-web/internal/pkg/util"
	"go-web/pkg/model"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/*
	使用gin-jwt中间件，只需要设置并实例化jwt.GinJWTMiddleware即可
	官方参考地址：https://github.com/appleboy/gin-jwt
*/

// login为登录处理函数，因为gin-jwt是授权认证一体的，不需要授权功能传入nil即可
func InitGinJWTMiddleware(login func(c *gin.Context) (interface{}, error)) (*jwt.GinJWTMiddleware, error) {

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           global.Conf.Jwt.Realm,                                 // jwt标识
		Key:             []byte(global.Conf.Jwt.Key),                           // 服务端密钥
		Timeout:         time.Hour * time.Duration(global.Conf.Jwt.Timeout),    // token过期时间
		MaxRefresh:      time.Hour * time.Duration(global.Conf.Jwt.MaxRefresh), // token最大刷新时间(RefreshToken过期时间=Timeout+MaxRefresh)
		PayloadFunc:     payloadFunc,                                           // 有效荷载处理
		IdentityHandler: identityHandler,                                       // 解析Claims
		Authenticator:   login,                                                 // 登录处理
		Authorizator:    authorizator,                                          // token校验成功处理
		Unauthorized:    unauthorized,                                          // token校验失败处理
		LoginResponse:   loginResponse,                                         // 登录成功后的响应
		LogoutResponse:  logoutResponse,                                        // 登出后的响应
		RefreshResponse: refreshResponse,                                       // 刷新token后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",    // 依次在这几个地方寻找请求中的token
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}

func payloadFunc(data interface{}) jwt.MapClaims {
	// 解析荷载
	if v, ok := data.(map[string]interface{}); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v["user"],
			"user":          v["user"],
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims[jwt.IdentityKey],
		"user":        claims["user"],
	}
}

// 认证成功处理
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		if user, ok := v["user"].(string); ok {
			userId := util.Str2Uint64(user)
			// 把userId设置到上下文中
			c.Set("user", userId)
			return true
		}
	}
	return false
}

// 认证失败处理
func unauthorized(c *gin.Context, code int, message string) {
	global.Log.Debug(c, "authorized fail...")
	response.FailWithCode(code)
}

func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	global.Log.Info(c, "login success...")
	response.SuccessWithData(map[string]interface{}{
		"token": token,
		"expires": model.LocalTime{
			Time: expires,
		},
	})
}

func logoutResponse(c *gin.Context, code int) {
	global.Log.Info(c, "logout success...")
	response.Success()
}

func refreshResponse(c *gin.Context, code int, token string, expires time.Time) {
	global.Log.Info(c, "refresh token success...")
	response.SuccessWithData(map[string]interface{}{
		"token": token,
		"expires": model.LocalTime{
			Time: expires,
		},
	})
}
