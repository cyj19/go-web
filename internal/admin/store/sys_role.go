package store

import (
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type SysRoleStore interface {
	UpdateMenuForRole(cd *model.CreateDelete) error
	GetByName(name string) (*model.SysRole, error)
}

type role struct {
	db *gorm.DB
}

func newSysRole(ds *datastore) SysRoleStore {
	return &role{db: ds.db}
}

//实现store.RoleStore接口

// 更新角色菜单(添加、删除)
func (r *role) UpdateMenuForRole(cd *model.CreateDelete) error {
	role := new(model.SysRole)
	role.Id = cd.Id
	var err error
	// 开启事务
	tx := r.db.Begin()
	if len(cd.Delete) > 0 {
		deleteMenus := make([]model.SysMenu, 0)
		for _, v := range cd.Delete {
			deleteMenus = append(deleteMenus, model.SysMenu{Model: model.Model{Id: v}})
		}
		// 删除关联
		err = tx.Model(role).Association("Menus").Delete(deleteMenus)
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return err
		}
	}

	if len(cd.Create) > 0 {
		createMenus := make([]model.SysMenu, 0)
		for _, v := range cd.Create {
			createMenus = append(createMenus, model.SysMenu{Model: model.Model{Id: v}})
		}
		// 添加关联
		err = tx.Model(role).Association("Menus").Append(createMenus)
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return err
		}
	}
	// 提交事务
	tx.Commit()
	return nil

}

func (r *role) GetByName(name string) (*model.SysRole, error) {
	result := &model.SysRole{}
	err := r.db.Preload("Menus").Where("name = ?", name).Order("sort").First(result).Error
	return result, err
}
