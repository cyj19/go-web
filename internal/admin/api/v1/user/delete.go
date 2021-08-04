package user

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (u *SysUserHandler) Delete(c *gin.Context) {
	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, 405, err, nil)
		return
	}

	ids := util.Str2Uint64Array(param.Ids)
	err = u.srv.SysUser().DeleteBatch(ids)
	if err != nil {
		util.WriteResponse(c, 500, err, nil)
		return
	}

	util.WriteResponse(c, 200, nil, nil)
}
