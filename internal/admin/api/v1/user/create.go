package user

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

//增加用户
func (u *SysUserHandler) Create(c *gin.Context) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	param.Password = util.EncryptionPsw(param.Password)
	err = u.srv.SysUser().Create(param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
