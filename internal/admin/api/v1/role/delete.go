package role

import (
	"strconv"

	"go-web/internal/admin/common"
	"go-web/internal/pkg/util"
	"go-web/pkg/errors"

	"github.com/gin-gonic/gin"
)

/*
	DELETE: /v1/role/:id
*/
func (r *RoleHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := r.srv.SysRole().Delete(uint64(id))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete role"), nil)
		return
	}
	//删除casbin_rule中的关联数据
	err = common.CasbinDeleteRole(c.Param("id"))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete role"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}

/*
	DELETE:  /v1/role?ids=1&ids=2&ids=3
*/
func (r *RoleHandler) DeleteBatch(c *gin.Context) {

	strs := c.QueryArray("ids")
	ids, err := common.ConverSliceToUint64(strs)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to conver param"), nil)
		return
	}
	err = r.srv.SysRole().DeleteBatch(ids)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete role collection"), nil)
		return
	}
	//删除casbin_rule中的关联数据
	err = common.CasbinDeleteRole(strs...)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to delete data in casbin_rule"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)

}
