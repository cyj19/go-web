package menu

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *SysMenuHandler) Update(c *gin.Context) {
	var menu model.SysMenu
	err := c.ShouldBind(&menu)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	err = m.srv.SysMenu().Update(c, &menu)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(menu)
}
