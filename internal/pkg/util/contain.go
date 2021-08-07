package util

import "go-web/internal/pkg/model"

func Contains(arr interface{}, item interface{}) bool {
	// 判断arr类型
	switch v := arr.(type) {
	case []model.SysMenu:
		var menus []model.SysMenu
		var menu model.SysMenu
		menus = v
		menu = item.(model.SysMenu)
		return ContainsSysMenu(menus, menu)
	}
	return false
}

func ContainsSysMenu(arr []model.SysMenu, item model.SysMenu) bool {
	for _, menu := range arr {
		if menu.Id == item.Id || menu.Name == item.Name {
			return true
		}
	}
	return false
}
