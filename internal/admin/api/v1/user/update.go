package user

import (
	"github.com/vagaryer/go-web/internal/pkg/model"
	"github.com/vagaryer/go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (u *SysUserHandler) Update(c *gin.Context) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	err = u.srv.SysUser().Update(c, &param)
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

	err = u.srv.SysUser().UpdateRoleForUser(c, &param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
