package menu

import (
	"errors"
	"strconv"

	"go-web/internal/admin/common"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (m *MenuHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := m.srv.SysMenu().Delete(uint64(id))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete menu"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}

/*
	DELETE: /v1/menu?ids=1&ids=2&ids=3
*/
func (m *MenuHandler) DeleteBatch(c *gin.Context) {

	strs := c.QueryArray("ids")
	ids, err := common.ConverSliceToUint64(strs)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete menus"), nil)
		return
	}
	err = m.srv.SysMenu().DeleteBatch(ids)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete menus"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
