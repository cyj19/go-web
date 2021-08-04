package api

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

func (a *SysApiHandler) Create(c *gin.Context) {
	var param model.SysApi
	err := c.ShouldBind(&param)
	if err != nil {
		log.Fatalf("参数绑定失败：%v", err)
		return
	}

	err = a.srv.Create(&param)
	if err != nil {
		log.Fatalf("增加接口失败：%v", err)
		return
	}
	util.WriteResponse(c, 200, nil, param)
}
