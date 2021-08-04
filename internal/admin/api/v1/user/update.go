package user

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

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
