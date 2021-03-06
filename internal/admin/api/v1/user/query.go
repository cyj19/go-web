package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cyj19/go-web/internal/pkg/model"
	"github.com/cyj19/go-web/internal/pkg/response"
	"github.com/cyj19/go-web/internal/pkg/util"
)

func (u *SysUserHandler) GetUserInfo(c *gin.Context) {
	currentUser := u.GetCurrentUser(c)
	var userResp model.SysUserResponse
	util.Struct2Struct(currentUser, &userResp)
	response.SuccessWithData(userResp)
}

func (u *SysUserHandler) GetCurrentUser(c *gin.Context) model.SysUser {
	userId, exist := c.Get("user")
	var currentUser *model.SysUser
	if !exist {
		return *currentUser
	}
	// 查询用户
	currentUser, _ = u.srv.SysUser().GetById(c, userId.(uint64))
	return *currentUser
}

//查询
func (u *SysUserHandler) GetByUsername(c *gin.Context) {

	user, err := u.srv.SysUser().GetByUsername(c, c.Param("name"))
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	var userResp model.SysUserResponse
	util.Struct2Struct(user, &userResp)
	response.SuccessWithData(userResp)

}

//查询多条记录，参数为json格式
func (u *SysUserHandler) GetList(c *gin.Context) {
	var param model.SysUser
	// 此处不能传入空指针，否则绑定失败
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}

	list, err := u.srv.SysUser().GetList(c, param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	var userRespList []model.SysUserResponse
	util.Struct2Struct(list, &userRespList)
	response.SuccessWithData(userRespList)
}

func (u *SysUserHandler) GetPage(c *gin.Context) {
	var param model.SysUserPage
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	fmt.Printf("userPage: %+v \n", param)
	page, err := u.srv.SysUser().GetPage(c, param)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.SuccessWithData(page)
}

// 使用go-jwt授权
func (u *SysUserHandler) Login(c *gin.Context) (interface{}, error) {
	var param model.SysUser
	err := c.ShouldBind(&param)
	if err != nil {
		return nil, err
	}

	user, err := u.srv.SysUser().Login(c, param.Username, param.Password)

	if err != nil || user == nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": fmt.Sprintf("%d", user.Id),
	}, nil
}
