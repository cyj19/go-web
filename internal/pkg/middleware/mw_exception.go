package middleware

import (
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			switch v := err.(type) {
			case response.Result:
				// 写入日志文件
				global.Log.Infof("response success, data: %v", v)
				response.JSON(c, v.Code, v)
				return
			default:
				// 写入日志文件
				global.Log.Errorf("unknown exception: %v", err)
				result := response.Result{
					Code: response.InternalServerError,
					Msg:  response.CustomError[response.InternalServerError],
					Data: nil,
				}
				response.JSON(c, result.Code, result)
				c.Abort()
				return
			}

		}
	}()
	c.Next()
}
