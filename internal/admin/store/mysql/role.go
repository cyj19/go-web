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
	return r.db.Save(role).Error
}

func (r *role) Delete(id uint64) error {
	return delete(r.db, id, &model.SysRole{})

}

func (r *role) DeleteBatch(ids []uint64) error {
	return deleteBatch(r.db, ids, &model.SysRole{})
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

func (r *role) List(whereOrder ...model.WhereOrder) ([]model.SysRole, error) {
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
