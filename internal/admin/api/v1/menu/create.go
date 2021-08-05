package menu

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (m *SysMenuHandler) Create(c *gin.Context) {
	var menu model.SysMenu
	err := c.ShouldBind(&menu)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}

	err = m.srv.Create(&menu)
	if err != nil {
		log.Fatalf("创建菜单失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, menu)
}
