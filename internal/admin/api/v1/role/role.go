package role

import (
	"go-web/internal/admin/common"
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewRoleHandler(factory store.Factory) *RoleHandler {
	return &RoleHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}

/*
	POST: /v1/role/permission
	Content-Type: x-www-form-urlencoded
*/
func (r *RoleHandler) SetRolePermission(c *gin.Context) {
	roleid := c.Query("rid")
	menuids := c.QueryArray("mids")

	err := common.CasbinSetRolePermission(r.srv, roleid, menuids...)
	if err != nil {
		util.WriteResponse(c, 500, err, nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
