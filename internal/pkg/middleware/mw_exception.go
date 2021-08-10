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
				global.LoggerIns.Println("正常响应：", err)
				response.JSON(c, v.Code, v)
				return
			default:
				// 写入日志文件
				global.LoggerIns.Println("未知异常：", err)
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
