package menu

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *SysMenuHandler) Create(c *gin.Context) {
	var menu model.SysMenu
	err := c.ShouldBind(&menu)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	err = m.srv.Create(&menu)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}
