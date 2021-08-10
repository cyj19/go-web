package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) Update(c *gin.Context) {
	var param model.SysApi
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	err = a.srv.SysApi().Update(&param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
