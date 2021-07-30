package menu

import (
	"errors"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (m *MenuHandler) Update(c *gin.Context) {
	var menu model.SysMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	err = m.srv.SysMenu().Update(&menu)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to update menu"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, menu)
}