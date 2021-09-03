package middleware

import (
	"github.com/vagaryer/go-web/internal/pkg/logger"
	"github.com/vagaryer/go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func Exception(glog *logger.GormZapLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				switch v := err.(type) {
				case response.Result:
					// 写入日志文件
					glog.Info(c, "response success, data: %v", v)
					response.JSON(c, v.Code, v)
					return
				default:
					// 写入日志文件
					glog.Error(c, "unknown exception: %v", err)
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

}
