package role

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"go-web/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (r *RoleHandler) Create(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBindJSON(&role)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	err = r.srv.Create(&role)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to create role"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, role)

}
