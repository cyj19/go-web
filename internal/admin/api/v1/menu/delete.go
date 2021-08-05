package menu

import (
	"errors"

	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

/*
	DELETE: /v1/menu/delete
*/
func (m *SysMenuHandler) BatchDelete(c *gin.Context) {

	strs := c.QueryArray("ids")
	ids, err := util.ConverSliceToUint64(strs)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete menus"), nil)
		return
	}
	err = m.srv.SysMenu().BatchDelete(ids)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete menus"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
