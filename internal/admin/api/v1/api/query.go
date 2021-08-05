package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) GetList(c *gin.Context) {
	var param model.SysApi
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	whereOrders := createSysApiQueryCondition(param)
	apis, err := a.srv.SysApi().GetList(whereOrders...)
	if err != nil {
		log.Fatalf("查询失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, apis)
}

func (a *SysApiHandler) GetPage(c *gin.Context) {
	var param model.SysApiPage
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	whereOrders := createSysApiQueryCondition(param.SysApi)
	apis, count, err := a.srv.SysApi().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		log.Fatalf("查询失败：%v", err)
		return
	}
	page := model.Page{
		Records: apis,
		Total:   count,
		PageInfo: model.PageInfo{
			PageIndex: param.PageIndex,
			PageSize:  param.PageSize,
		},
	}
	page.SetPageNum(count)
	util.WriteResponse(c, 200, nil, page)
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
