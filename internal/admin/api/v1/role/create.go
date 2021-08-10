package role

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) Create(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBind(&role)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	err = r.srv.Create(&role)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()

}
