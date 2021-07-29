package user

import (
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func LoadPolicy(c *gin.Context) {
	enforcer := initialize.GetEnforcerIns()
	err := enforcer.LoadPolicy()
	if err != nil {
		util.WriteResponse(c, 500, err, nil)
		return
	}
	util.WriteResponse(c, 200, nil, nil)
}
