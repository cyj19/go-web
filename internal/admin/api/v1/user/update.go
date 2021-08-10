package user

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (u *SysUserHandler) Update(c *gin.Context) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	err = u.srv.SysUser().Update(&param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}

func (u *SysUserHandler) UpdateRoleForUser(c *gin.Context) {
	var param model.CreateDelete
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = u.srv.SysUser().UpdateRoleForUser(&param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
