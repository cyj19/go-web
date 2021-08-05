package menu

import (
	"errors"
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
	err := c.ShouldBindJSON(&param)
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

func (m *SysMenuHandler) GetPage(c *gin.Context) {
	var param model.SysMenuPage
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	whereOrders := createSysMenuQueryCondition(param.SysMenu)
	menus, count, err := m.srv.SysMenu().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get menu page"), nil)
		return
	}
	page := &model.Page{
		Records:  menus,
		PageInfo: model.PageInfo{PageIndex: param.PageIndex, PageSize: param.PageSize},
	}
	page.SetPageNum(count)
	util.WriteResponse(c, 200, nil, page)
}

func createSysMenuQueryCondition(param model.SysMenu) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)

	if param.Name != "" {
		v := "%" + param.Name + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
	}

	whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{param.Status}})

	return whereOrders
}
