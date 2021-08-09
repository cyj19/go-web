package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
)

func (u *SysUserHandler) GetUserInfo(c *gin.Context) {
	currentUser := u.GetCurrentUser(c)
	var userResp model.SysUserResponse
	util.Struct2Struct(currentUser, &userResp)
	util.WriteResponse(c, 0, nil, userResp)
}

func (u *SysUserHandler) GetCurrentUser(c *gin.Context) model.SysUser {
	userId, exist := c.Get("user")
	var currentUser *model.SysUser
	if !exist {
		return *currentUser
	}
	// 查询用户
	currentUser, _ = u.srv.SysUser().GetById(userId.(uint64))
	return *currentUser
}

//查询
func (u *SysUserHandler) GetByUsername(c *gin.Context) {

	user, err := u.srv.SysUser().GetByUsername(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	var userResp model.SysUserResponse
	util.Struct2Struct(user, &userResp)
	util.WriteResponse(c, 0, nil, userResp)

}

//查询多条记录，参数为json格式
func (u *SysUserHandler) GetList(c *gin.Context) {
	var param model.SysUser
	// 此处不能传入空指针，否则绑定失败
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	whereOrders := createSysUserQueryCondition(param)
	list, err := u.srv.SysUser().GetList(whereOrders...)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	var userRespList []model.SysUserResponse
	util.Struct2Struct(list, &userRespList)
	util.WriteResponse(c, 0, nil, userRespList)
}

func (u *SysUserHandler) GetPage(c *gin.Context) {
	var param model.SysUserPage
	err := c.ShouldBind(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	whereOrders := createSysUserQueryCondition(param.SysUser)
	page, err := u.srv.SysUser().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	util.WriteResponse(c, 0, nil, page)
}

// 使用go-jwt授权
func (u *SysUserHandler) Login(c *gin.Context) (interface{}, error) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		return nil, err
	}

	user, err := u.srv.SysUser().Login(param.Username, param.Password)

	if err != nil || user == nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": fmt.Sprintf("%d", user.Id),
	}, nil
}

func createSysUserQueryCondition(param model.SysUser) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)

	if param.Id > 0 {
		v := param.Id
		whereOrders = append(whereOrders, model.WhereOrder{Where: "id = ?", Value: []interface{}{v}})
	}
	if param.Username != "" {
		v := "%" + param.Username + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "username like ?", Value: []interface{}{v}})
	}
	if param.Status != nil {
		whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*param.Status}})
	}

	return whereOrders
}
