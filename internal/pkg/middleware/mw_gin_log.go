package middleware

import (
	"fmt"
	"go-web/internal/pkg/global"
	"time"

	"github.com/gin-gonic/gin"
)

// zap替换gin的默认日志
func GinLog(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	// 请求IP
	clientIP := c.ClientIP()
	// 请求方式
	method := c.Request.Method
	// 请求路径
	path := c.Request.URL.Path
	// 请求参数
	query := c.Request.URL.RawQuery

	// 处理请求
	c.Next()
	// 响应状态码
	statusCode := c.Writer.Status()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	execTime := endTime.Sub(startTime)

	global.Log.Info(fmt.Sprintf("%d %s %s %s %s %s", statusCode, method, path, query, clientIP, execTime))

}
