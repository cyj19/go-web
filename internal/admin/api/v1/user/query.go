package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
)

//查询
func (u *UserHandler) GetByUsername(c *gin.Context) {

	user, err := u.srv.SysUser().GetByUsername(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	util.WriteResponse(c, 0, nil, user)

}

//查询多条记录，参数为json格式
func (u *UserHandler) List(c *gin.Context) {
	var param *model.SysUser
	err := c.ShouldBindJSON(param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	list, err := u.srv.SysUser().List(param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	util.WriteResponse(c, 0, nil, list)
}

func (u *UserHandler) GetPage(c *gin.Context) {
	var param *model.SysUserPage
	err := c.ShouldBindJSON(param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	list, count, err := u.srv.SysUser().GetPage(param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	page := &model.Page{
		Records:  list,
		PageInfo: model.PageInfo{PageIndex: param.PageIndex, PageSize: param.PageSize},
	}
	page.SetPageNum(count)
	util.WriteResponse(c, 0, nil, page)
}
