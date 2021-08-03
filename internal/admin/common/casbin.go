package common

import (
	"errors"

	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/util"
)

//删除角色及关联数据
func CasbinDeleteRole(roleids ...string) error {
	enforcer := initialize.GetEnforcerIns()
	for _, rid := range roleids {
		enforcer.DeletePermissionsForUser(rid)
		enforcer.DeleteRole(rid)
	}
	return nil
}

//设置角色权限 (先删除，后添加)
func CasbinSetRolePermission(srv srvv1.Service, roleid string, menuids ...string) error {
	enforcer := initialize.GetEnforcerIns()

	//查询菜单
	ids, _ := util.ConverSliceToUint64(menuids)
	menus, err := srv.SysMenu().GetSome(ids)
	if len(menus) < 1 || err != nil {
		return errors.New("failed to get menus by ids")
	}

	enforcer.DeletePermissionsForUser(roleid)

	// for _, menu := range menus {
	// 	if menu.Type == 3 {
	// 		enforcer.AddPermissionForUser(roleid, menu.URL, menu.Method)
	// 	}

	// }
	return nil
}

//查询角色权限
func CasbinGetRolePermission(roleid string) [][]string {
	enforcer := initialize.GetEnforcerIns()
	return enforcer.GetPermissionsForUser(roleid)
}

//用户设置角色 (先删除，后添加)
func CasbinSetUserRole(srv srvv1.Service, userid string, roleids ...string) error {
	enforcer := initialize.GetEnforcerIns()
	enforcer.DeleteRolesForUser(userid)
	_, err := enforcer.AddRolesForUser(userid, roleids)
	return err
}
