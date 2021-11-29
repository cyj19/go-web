package role

import (
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) Create(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBind(&role)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	err = r.srv.SysRole().Create(c, role)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()

}
