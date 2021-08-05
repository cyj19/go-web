package menu

import (
	"log"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

/*
	DELETE: /v1/menu/delete
*/
func (m *SysMenuHandler) BatchDelete(c *gin.Context) {

	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	ids := util.Str2Uint64Array(param.Ids)
	err = m.srv.SysMenu().BatchDelete(ids)
	if err != nil {
		log.Fatalf("删除菜单失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
