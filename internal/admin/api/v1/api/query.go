package api

import (
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) GetList(c *gin.Context) {
	var param model.SysApi
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	apis, err := a.srv.SysApi().GetList(c, param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(apis)
}

func (a *SysApiHandler) GetPage(c *gin.Context) {
	var param model.SysApiPage
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	page, err := a.srv.SysApi().GetPage(c, param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}
