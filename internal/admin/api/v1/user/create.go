package user

import (
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"
	"github.com/cyj19/go-web/internal/pkg/util"

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
	err = u.srv.SysUser().Create(c, param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
