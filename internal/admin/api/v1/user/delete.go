package user

import (
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"
	"github.com/cyj19/go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (u *SysUserHandler) BatchDelete(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	ids := util.Str2Uint64Array(param.Ids)
	err = u.srv.SysUser().BatchDelete(c, ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
