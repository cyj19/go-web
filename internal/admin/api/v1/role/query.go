package role

import (
	"errors"
	"strconv"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (r *RoleHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := r.srv.SysRole().GetById(uint64(id))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get role"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, role)
}

func (r *RoleHandler) List(c *gin.Context) {
	var param model.SysRole
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}

	roles, err := r.srv.SysRole().List(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get roles"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, roles)
}

func (r *RoleHandler) GetPage(c *gin.Context) {
	var param model.SysRolePage
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}

	roles, count, err := r.srv.SysRole().GetPage(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get role page"), nil)
		return
	}

	page := &model.Page{
		Records:  roles,
		PageInfo: model.PageInfo{PageIndex: param.PageIndex, PageSize: param.PageSize},
	}
	page.SetPageNum(count)
	util.WriteResponse(c, 200, nil, page)
}
