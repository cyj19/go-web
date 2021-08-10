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
	whereOrders := createSysRoleQueryCondition(param)
	roles, err := r.srv.SysRole().GetList(whereOrders...)
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
	whereOrders := createSysRoleQueryCondition(param.SysRole)
	page, err := r.srv.SysRole().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}

func createSysRoleQueryCondition(param model.SysRole) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)

	if param.Name != "" {
		v := "%" + param.Name + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
	}
	if param.NameZh != "" {
		v := "%" + param.NameZh + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "name_zh like ?", Value: []interface{}{v}})
	}
	if param.Status != nil {
		whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*param.Status}})
	}
	if param.Sort != nil {
		whereOrders = append(whereOrders, model.WhereOrder{Where: "sort = ?", Value: []interface{}{*param.Sort}})
	}
	whereOrders = append(whereOrders, model.WhereOrder{Order: "sort"})

	return whereOrders
}
