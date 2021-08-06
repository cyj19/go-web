package admin

import (
	"fmt"
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
	"strings"

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
	apis := mockSysApi(configuration.Server.ApiVersion)
	newApis := make([]model.SysApi, 0)
	newRoleCasbins := make([]model.SysRoleCasbin, 0)
	for _, api := range apis {
		whereOrder := model.WhereOrder{Where: "method = ? and path = ?", Value: []interface{}{api.Method, api.Path}}
		oldApis, err := service.SysApi().GetList(whereOrder)
		if err != nil || len(oldApis) == 0 {
			newApis = append(newApis, api)
			p := strings.TrimPrefix(api.Path, "/"+configuration.Server.ApiVersion)

			// 不需要权限验证的接口不加入casbin规则
			basePaths := map[string]string{
				"/base/login":         "",
				"/base/logout":        "",
				"/base/refresh_token": "",
			}

			if _, ok := basePaths[p]; ok {
				continue
			}
			// 构建casbin规则
			// 管理员拥有所有接口权限
			newRoleCasbins = append(newRoleCasbins, model.SysRoleCasbin{
				Kyeword: util.Uint642Str(roles[0].Id),
				Path:    api.Path,
				Method:  api.Method,
			})

			// 公共接口
			publicPaths := map[string]string{
				"/user/info": "",
			}

			if _, ok := publicPaths[p]; ok {
				for i := 1; i < len(roles); i++ {
					// 其他角色赋予基础接口权限
					newRoleCasbins = append(newRoleCasbins, model.SysRoleCasbin{
						Kyeword: util.Uint642Str(roles[i].Id),
						Path:    api.Path,
						Method:  api.Method,
					})
				}

			}
		}
	}
	if len(newApis) > 0 {
		err := service.Create(&newApis)
		if err != nil {
			panic(fmt.Errorf("初始化接口失败：%v", err))
		}
	}
	if len(newRoleCasbins) > 0 {
		_, err := service.SysCasbin().BatchCreateRoleCasbins(newRoleCasbins)
		if err != nil {
			panic(fmt.Errorf("初始化接口权限失败：%v", err))
		}
	}

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
func mockSysApi(version string) []model.SysApi {
	apiVersion := "/" + version
	return []model.SysApi{
		{
			Method:   "POST",
			Path:     apiVersion + "/base/login",
			Category: "base",
			Creator:  "系统创建",
		},
		{
			Method:   "GET",
			Path:     apiVersion + "/base/logout",
			Category: "base",
			Creator:  "系统创建",
		},
		{
			Method:   "GET",
			Path:     apiVersion + "/base/refresh_token",
			Category: "base",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/user/add",
			Category: "user",
			Creator:  "系统创建",
		},
		{
			Method:   "DELETE",
			Path:     apiVersion + "/user/delete",
			Category: "user",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/user/update",
			Category: "user",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/user/role/update",
			Category: "user",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/user/page",
			Category: "user",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/role/add",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "DELETE",
			Path:     apiVersion + "/role/delete",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/role/update",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/role/menu/update",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/role/api/update",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/role/list",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/role/page",
			Category: "role",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/menu/add",
			Category: "menu",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/menu/update",
			Category: "menu",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/menu/list",
			Category: "menu",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/menu/page",
			Category: "menu",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/api/add",
			Category: "api",
			Creator:  "系统创建",
		},
		{
			Method:   "DELETE",
			Path:     apiVersion + "/api/delete",
			Category: "api",
			Creator:  "系统创建",
		},
		{
			Method:   "PATCH",
			Path:     apiVersion + "/api/update",
			Category: "api",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/api/list",
			Category: "api",
			Creator:  "系统创建",
		},
		{
			Method:   "POST",
			Path:     apiVersion + "/api/page",
			Category: "api",
			Creator:  "系统创建",
		},
	}
}
