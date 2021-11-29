package util

import (
	"fmt"
	"testing"

	"github.com/cyj19/go-web/internal/pkg/model"
)

func TestStruct2Struct2(t *testing.T) {
	status := true
	roles := make([]model.SysRole, 0)
	roles = append(roles, model.SysRole{
		Name:   "roletest",
		NameZh: "测试员",
	})
	user := model.SysUser{
		Username: "test",
		Password: "123456",
		Status:   &status,
		Roles:    roles,
	}
	var userResp model.SysUserResponse
	Struct2Struct(user, &userResp)
	fmt.Printf("userResp: %+v \n", userResp)
}

func TestList2List(t *testing.T) {
	roles := make([]model.SysRole, 0)
	roles = append(roles, model.SysRole{
		Name:   "roletest",
		NameZh: "测试员",
	})
	users := make([]model.SysUser, 0)
	users = append(users, model.SysUser{
		Username: "test",
		Password: "123456",
		Roles:    roles,
	})
	var userRespList []model.SysUserResponse
	Struct2Struct(users, &userRespList)
	fmt.Printf("userRespList: %+v \n", userRespList)
}
