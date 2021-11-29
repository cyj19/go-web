package role

import (
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"
	"github.com/cyj19/go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) BatchDelete(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	ids := util.Str2Uint64Array(param.Ids)
	err = r.srv.SysRole().BatchDelete(c, ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
