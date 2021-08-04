package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) GetPage(c *gin.Context) {
	var param model.SysApiPage
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	apis, count, err := a.srv.SysApi().GetPage(&param)
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
