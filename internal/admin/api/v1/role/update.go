package role

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"go-web/pkg/errors"
	"log"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) Update(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBindJSON(&role)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
	}
	err = r.srv.SysRole().Update(&role)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to update"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, role)
}

func (r *SysRoleHandler) UpdateMenuForRole(c *gin.Context) {
	var param model.CreateDelete
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}

	err = r.srv.SysRole().UpdateMenuForRole(&param)
	if err != nil {
		log.Fatalf("更新失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, "")
}
