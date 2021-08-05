package admin

import (
	"fmt"
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"

	"github.com/casbin/casbin/v2"
)

// 初始化数据
func InitData(factoryIns store.Factory, enforcer *casbin.Enforcer) {
	configuration := initialize.GetConfiguration()
	if !configuration.Server.InitData {
		return
	}

	service := srvv1.NewService(factoryIns, enforcer)

	// 初始化角色
	newRoles := make([]model.SysRole, 0)

	roles := []model.SysRole{
		{
			Name:   "admin",
			NameZh: "管理员",
		},
		{
			Name:   "guest",
			NameZh: "访客",
		},
	}

	for i, value := range roles {
		oldRole, err := service.SysRole().GetByName(value.Name)
		if err != nil || oldRole == nil {
			newRoles = append(newRoles, value)
		} else {
			roles[i].Id = oldRole.Id
		}
	}

	if len(newRoles) > 0 {
		service.Create(&newRoles)
		// 如果admin 和 guest都插入
		if len(newRoles) == len(roles) {
			roles = newRoles
		} else {
			// 只插入一个角色，赋值到roles相应的元素中
			for i := range roles {
				if roles[i].Name == newRoles[0].Name {
					roles[i] = newRoles[0]
				}
			}
		}
	}

	// 初始化菜单

	menus := []model.SysMenu{
		{
			Name:  "dashboardRoot",
			Title: "首页父菜单",
			Icon:  "dashboard",
			Path:  "/dashboard",
			Roles: roles,
			Children: []model.SysMenu{
				{
					Name:      "dashboard",
					Title:     "首页",
					Icon:      "dashboard",
					Path:      "index",
					Component: "/dashboard/index",
					Roles:     roles,
				},
			},
		},
		{
			Name:  "systemRoot",
			Title: "系统设置",
			Icon:  "component",
			Path:  "/system",
			Children: []model.SysMenu{
				{
					Name:      "menu",
					Title:     "菜单管理",
					Icon:      "tree-table",
					Path:      "menu", // 前端中子菜单对继承父菜单的路径
					Component: "/system/menu",
				},
				{
					Name:      "role",
					Title:     "角色管理",
					Icon:      "peoples",
					Path:      "role", // 前端中子菜单对继承父菜单的路径
					Component: "/system/role",
				},
				{
					Name:      "user",
					Title:     "用户管理",
					Icon:      "user",
					Path:      "user", // 前端中子菜单对继承父菜单的路径
					Component: "/system/user",
				},
				{
					Name:      "api",
					Title:     "接口管理",
					Icon:      "tree",
					Path:      "api", // 前端中子菜单对继承父菜单的路径
					Component: "/system/api",
				},
			},
		},
	}
	// 生成菜单，先创建父菜单，再创建子菜单
	generateMenu(0, menus, roles[0], service)

	// 初始化接口
	apis := mockSysApi()
	service.Create(apis)

	// 初始化用户
	newUsers := make([]model.SysUser, 0)

	users := []model.SysUser{
		{
			Username: "admin",
			Password: "123456",
			Roles:    []model.SysRole{roles[0]}, // 默认拥有admin角色
		},
		{
			Username: "guest",
			Password: "123456",
			Roles:    []model.SysRole{roles[1]}, // 默认拥有guest角色
		},
	}

	for _, value := range users {
		oldUser, err := service.SysUser().GetByUsername(value.Username)
		if err != nil || oldUser == nil {
			newUsers = append(newUsers, value)
		}
	}

	if len(newUsers) > 0 {
		service.Create(&newUsers)
	}

}

func generateMenu(parentId uint64, menus []model.SysMenu, adminRole model.SysRole, srv srvv1.Service) {

	if len(menus) > 0 {
		newMenus := make([]model.SysMenu, 0)

		// 创建父菜单
		for i, value := range menus {
			value.ParentId = parentId
			oldMenu, err := srv.SysMenu().GetByPath(value.Path)
			if err != nil || oldMenu == nil {
				sort := uint(i)
				value.Sort = &sort
				if len(value.Roles) == 0 {
					value.Roles = []model.SysRole{adminRole}
				}
				newMenus = append(newMenus, value)
			}

		}
		if len(newMenus) > 0 {
			err := srv.Create(&newMenus)
			if err != nil {
				panic(fmt.Errorf("初始化菜单失败：%v", err))
			}

			// 创建子菜单
			for i := range newMenus {
				value := newMenus[i]
				if len(value.Children) > 0 {

					for j := range value.Children {
						sort := uint(j)
						// 添加菜单顺序，因为要修改值，所以要用索引来获取
						value.Children[j].Sort = &sort
						// 添加父菜单id，因为要修改值，所以要用索引来获取
						value.Children[j].ParentId = value.Id
						if len(value.Children[j].Roles) == 0 {
							value.Children[j].Roles = []model.SysRole{adminRole}
						}
					}
				}
				if len(value.Roles) == 0 {
					value.Roles = []model.SysRole{adminRole}
				}
				err := srv.Create(&value.Children)
				if err != nil {
					panic(fmt.Errorf("初始化子菜单失败：%v", err))
				}

			}
		}

	}

}

// 初始化的接口数据
func mockSysApi() []model.SysApi {
	return []model.SysApi{
		{
			Method:   "post",
			Path:     "/v1/base/login",
			Category: "base",
			Creator:  "系统创建",
		},
	}
}
