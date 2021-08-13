package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) BatchDelete(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	ids := util.Str2Uint64Array(param.Ids)
	err = a.srv.SysApi().BatchDelete(ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
