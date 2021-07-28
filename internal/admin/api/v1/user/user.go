package user

import (
	"go-web/internal/admin/common"
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewUserHandler(factory store.Factory) *UserHandler {
	return &UserHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}

/*
	POST: /v1/user/role
	Content-Type: x-www-form-urlencoded
*/
func (u *UserHandler) SetUserRole(c *gin.Context) {
	userid := c.PostForm("uid")
	roleids := c.PostFormArray("rids")
	err := common.CasbinSetUserRole(u.srv, userid, roleids...)
	if err != nil {
		util.WriteResponse(c, 500, err, nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
