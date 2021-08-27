package role

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) Update(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBind(&role)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
	}
	err = r.srv.SysRole().Update(c, &role)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(role)
}

func (r *SysRoleHandler) UpdateMenuForRole(c *gin.Context) {
	var param model.CreateDelete
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = r.srv.SysRole().UpdateMenuForRole(c, &param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

func (r *SysRoleHandler) UpdateApiForRole(c *gin.Context) {
	var param model.CreateDelete
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = r.srv.SysRole().UpdateApiForRole(c, &param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
