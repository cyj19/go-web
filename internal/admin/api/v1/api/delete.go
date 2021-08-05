package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) BatchDelete(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	ids := util.Str2Uint64Array(param.Ids)
	err = a.srv.SysApi().BatchDelete(ids)
	if err != nil {
		log.Fatalf("删除接口失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
