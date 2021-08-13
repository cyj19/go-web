package role

import (
	"strconv"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) GetById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	role, err := r.srv.SysRole().GetById(uint64(id))
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.SuccessWithData(role)
}

func (r *SysRoleHandler) GetList(c *gin.Context) {
	var param model.SysRole
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	roles, err := r.srv.SysRole().GetList(param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(roles)
}

func (r *SysRoleHandler) GetPage(c *gin.Context) {
	var param model.SysRolePage
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	page, err := r.srv.SysRole().GetPage(param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}
