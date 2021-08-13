package menu

import (
	"strconv"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (m *SysMenuHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithCode(response.InternalServerError)
		return
	}
	menu, err := m.srv.SysMenu().GetById(uint64(id))
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(menu)
}

func (m *SysMenuHandler) GetList(c *gin.Context) {
	var param model.SysMenu
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	menus, err := m.srv.SysMenu().GetList(param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(menus)
}

func (m *SysMenuHandler) GetMenusByRoleId(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	menus, err := m.srv.SysMenu().GetMenusByRoleId(util.Str2Uint64Array(param.Ids))
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(menus)
}

func (m *SysMenuHandler) GetPage(c *gin.Context) {
	var param model.SysMenuPage
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	page, err := m.srv.SysMenu().GetPage(param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}
