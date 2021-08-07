package mysql

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type role struct {
	db *gorm.DB
}

func newSysRole(ds *datastore) store.SysRoleStore {
	return &role{db: ds.db}
}

//实现store.RoleStore接口

func (r *role) Create(role *model.SysRole) error {
	return r.db.Create(role).Error
}

func (r *role) Update(role *model.SysRole) error {
	return r.db.Updates(role).Error
}

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

func (r *role) BatchDelete(ids []uint64) error {
	return batchDelete(r.db, &model.SysRole{}, ids)
}

func (r *role) GetById(id uint64) (*model.SysRole, error) {
	result := &model.SysRole{}
	err := r.db.Preload("Menus").Where("id = ?", id).First(result).Error
	return result, err
}

func (r *role) GetByName(name string) (*model.SysRole, error) {
	result := &model.SysRole{}
	err := r.db.Preload("Menus").Where("name = ?", name).First(result).Error
	return result, err
}

func (r *role) GetList(whereOrder ...model.WhereOrder) ([]model.SysRole, error) {
	result := make([]model.SysRole, 0)
	tx := queryByCondition(r.db, &model.SysRole{}, whereOrder)
	err := tx.Preload("Menus").Find(&result).Error
	return result, err
}

func (r *role) GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysRole, int64, error) {
	result := make([]model.SysRole, 0)
	tx := queryByCondition(r.db, &model.SysRole{}, whereOrder)
	//查询总数
	var count int64
	var err error
	err = tx.Count(&count).Error
	//有错误或总数为0，直接返回
	if err != nil || count == 0 {
		return nil, count, err
	}
	err = tx.Preload("Menus").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&result).Error
	return result, count, err
}
