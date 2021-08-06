package role

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) Create(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBind(&role)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	err = r.srv.Create(&role)
	if err != nil {
		log.Fatalf("添加失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, role)

}
