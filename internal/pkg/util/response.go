package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteResponse(c *gin.Context, code int, err error, data interface{}) {
	if err != nil {
		c.JSON(code, gin.H{
			"code": code,
			"msg":  err.Error(),
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}
