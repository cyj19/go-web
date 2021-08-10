package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) GetList(c *gin.Context) {
	var param model.SysApi
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	whereOrders := createSysApiQueryCondition(param)
	apis, err := a.srv.SysApi().GetList(whereOrders...)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(apis)
}

func (a *SysApiHandler) GetPage(c *gin.Context) {
	var param model.SysApiPage
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	whereOrders := createSysApiQueryCondition(param.SysApi)
	page, err := a.srv.SysApi().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}

func createSysApiQueryCondition(param model.SysApi) []model.WhereOrder {

	whereOrders := make([]model.WhereOrder, 0)
	if param.Id > 0 {
		whereOrders = append(whereOrders, model.WhereOrder{Where: "id = ?", Value: []interface{}{param.Id}})
	}
	if param.Method != "" {
		v := "%" + param.Method + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "method like ?", Value: []interface{}{v}})
	}
	if param.Path != "" {
		v := "%" + param.Path + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "path like ?", Value: []interface{}{v}})
	}
	if param.Category != "" {
		v := "%" + param.Category + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "category like ?", Value: []interface{}{v}})
	}
	if param.Creator != "" {
		v := "%" + param.Creator + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "creator like ?", Value: []interface{}{v}})
	}
	return whereOrders

}
