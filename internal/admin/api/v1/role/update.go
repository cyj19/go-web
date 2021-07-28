package role

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"go-web/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (r *RoleHandler) Update(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBindJSON(&role)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
	}
	err = r.srv.SysRole().Update(&role)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to update"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, role)
}
