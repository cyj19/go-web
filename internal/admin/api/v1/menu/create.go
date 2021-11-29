package menu

import (
	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *SysMenuHandler) Create(c *gin.Context) {
	var menu model.SysMenu
	err := c.ShouldBind(&menu)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = m.srv.SysMenu().Create(c, menu)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
