package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) Create(c *gin.Context) {
	var param model.SysApi
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = a.srv.Create(&param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
