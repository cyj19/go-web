package menu

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"go-web/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (m *MenuHandler) Create(c *gin.Context) {
	var menu model.SysMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}

	menu.Status = true

	err = m.srv.SysMenu().Create(&menu)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to create menu"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, menu)
}
