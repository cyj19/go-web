package user

import (
	"net/http"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

//增加用户
func (u *UserHandler) Create(c *gin.Context) {
	var param model.SysUser
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	err = u.srv.Create(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	util.WriteResponse(c, 0, nil, param)
}
