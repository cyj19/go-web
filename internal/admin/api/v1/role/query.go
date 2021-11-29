package role

import (
	"strconv"

	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) GetById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	role, err := r.srv.SysRole().GetById(c, uint64(id))
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

	roles, err := r.srv.SysRole().GetList(c, param)
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

	page, err := r.srv.SysRole().GetPage(c, param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}
