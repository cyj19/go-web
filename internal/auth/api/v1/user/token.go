package user

import (
	"fmt"

	"go-web/internal/pkg/model"

	"github.com/gin-gonic/gin"
)

// 自定义jwt授权
// var jwtkey = []byte("go-web")

// func (u *UserHandler) Token(c *gin.Context) {
// 	var param model.SysUser
// 	//绑定参数
// 	err := c.ShouldBindJSON(&param)
// 	if err != nil {
// 		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
// 		return
// 	}
// 	var result *model.SysUser
// 	result, err = u.srv.User().Login(param.Username, param.Password)
// 	if err != nil {
// 		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
// 		return
// 	}

// 	//颁发token
// 	expireTime := time.Now().Add(2 * time.Hour)
// 	tokenString := productToken(c, result, expireTime)
// 	//写入redis
// 	redisIns := initialize.GetRedisIns()
// 	err = redisIns.Set(tokenString, "", 2*time.Hour).Err()

// 	if err != nil {
// 		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
// 		return
// 	}

// 	util.WriteResponse(c, 0, nil, tokenString)
// }

// func productToken(c *gin.Context, user *model.SysUser, expireTime time.Time) string {

// 	claims := &model.Claims{
// 		UserId: user.Id,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expireTime.Unix(), //过期时间
// 			IssuedAt:  time.Now().Unix(),
// 			Issuer:    "127.0.0.1",    //签名颁发者
// 			Subject:   "go-web token", //签名主题
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenStr, err := token.SignedString(jwtkey)
// 	if err != nil {
// 		return ""
// 	}
// 	return tokenStr
// }

// 使用go-jwt授权
func (u *UserHandler) Login(c *gin.Context) (interface{}, error) {
	var param model.SysUser
	err := c.ShouldBindJSON(&param)
	if err != nil {
		return nil, err
	}

	user, err := u.srv.User().Login(param.Username, param.Password)

	if err != nil || user == nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": fmt.Sprintf("%d", user.Id),
	}, nil
}
