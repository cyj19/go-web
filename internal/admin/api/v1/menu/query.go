package menu

import (
	"errors"
	"log"
	"strconv"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (m *SysMenuHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	menu, err := m.srv.SysMenu().GetById(uint64(id))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get menu"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, menu)
}

func (m *SysMenuHandler) GetList(c *gin.Context) {
	var param model.SysMenu
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	whereOrders := createSysMenuQueryCondition(param)
	menus, err := m.srv.SysMenu().GetList(whereOrders...)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get menu list"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, menus)
}

func (m *SysMenuHandler) GetMenusByRoleId(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	menus, err := m.srv.SysMenu().GetMenusByRoleId(util.Str2Uint64Array(param.Ids))
	if err != nil {
		log.Fatalf("查询失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, menus)
}

func (m *SysMenuHandler) GetPage(c *gin.Context) {
	var param model.SysMenuPage
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	whereOrders := createSysMenuQueryCondition(param.SysMenu)

	page, err := m.srv.SysMenu().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get menu page"), nil)
		return
	}

	util.WriteResponse(c, 200, nil, page)
}

func createSysMenuQueryCondition(param model.SysMenu) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)

	if param.Name != "" {
		v := "%" + param.Name + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
	}
	if param.Status != nil {
		whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*param.Status}})
	}

	whereOrders = append(whereOrders, model.WhereOrder{Order: "parent_id, sort"})

	return whereOrders
}
