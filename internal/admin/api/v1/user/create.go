package user

import (
	"net/http"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

//增加用户
func (u *SysUserHandler) Create(c *gin.Context) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	param.Password = util.EncryptionPsw(param.Password)
	err = u.srv.Create(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	util.WriteResponse(c, 0, nil, param)
}
