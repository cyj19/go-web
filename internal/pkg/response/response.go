package response

import "github.com/gin-gonic/gin"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WriteResult(code int, msg string, data interface{}) {
	// 主动抛出panic，由全局异常处理中间件来同一返回响应
	panic(Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success() {
	WriteResult(OK, CustomError[OK], nil)
}

func SuccessWithData(data interface{}) {
	WriteResult(OK, CustomError[OK], data)
}

func FailWithMsg(msg string) {
	if msg == "" {
		msg = CustomError[NotOK]
	}
	WriteResult(NotOK, msg, nil)
}

func FailWithCode(code int) {
	msg := CustomError[NotOK]
	if val, ok := CustomError[code]; ok {
		msg = val
	}
	WriteResult(code, msg, nil)
}

func FailWithCodeAndMsg(code int, msg string) {
	WriteResult(code, msg, nil)
}

func JSON(c *gin.Context, code int, result interface{}) {
	c.JSON(code, result)
}
