package user

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (u *SysUserHandler) Update(c *gin.Context) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}
	err = u.srv.SysUser().Update(&param)
	if err != nil {
		log.Fatalf("更新失败：%v", err)
		return
	}

	util.WriteResponse(c, 200, nil, "")
}

func (u *SysUserHandler) UpdateRoleForUser(c *gin.Context) {
	var param model.CreateDelete
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, 405, err, nil)
		return
	}

	err = u.srv.SysUser().UpdateRoleForUser(&param)
	if err != nil {
		util.WriteResponse(c, 500, err, nil)
		return
	}

	util.WriteResponse(c, 200, nil, "")
}
