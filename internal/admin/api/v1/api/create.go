package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) Create(c *gin.Context) {
	var api model.SysApi
	err := c.ShouldBind(&api)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = a.srv.SysApi().Create(api)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
