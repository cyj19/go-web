package mysql

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func newSysUser(ds *datastore) store.SysUserStore {
	return &user{db: ds.db}
}

//实现store.UserStore接口
// Updates使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段
func (u *user) Update(user *model.SysUser) error {
	return u.db.Updates(user).Error
}

// 更新用户角色(添加、删除)
func (u *user) UpdateRoleForUser(cd *model.CreateDelete) error {
	user := model.SysUser{}
	user.Id = cd.Id
	deleteRoles := make([]model.SysRole, 0)
	for _, v := range cd.Delete {
		deleteRoles = append(deleteRoles, model.SysRole{Model: model.Model{Id: v}})
	}

	createRoles := make([]model.SysRole, 0)
	for _, v := range cd.Create {
		createRoles = append(createRoles, model.SysRole{Model: model.Model{Id: v}})
	}
	// 开启事务
	tx := u.db.Begin()
	// 删除关联
	err := tx.Model(&user).Association("Roles").Delete(deleteRoles)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	// 添加关联
	err = tx.Model(&user).Association("Roles").Append(createRoles)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

func (u *user) BatchDelete(ids []uint64) error {
	return batchDelete(u.db, &model.SysUser{}, ids)
}

func (u *user) GetByUsername(username string) (*model.SysUser, error) {
	result := model.SysUser{}
	err := u.db.Preload("Roles").Where("username = ?", username).First(&result).Error
	return &result, err
}

// func (u *user) GetList(whereOrder ...model.WhereOrder) ([]model.SysUser, error) {
// 	result := make([]model.SysUser, 0)
// 	tx := queryByCondition(u.db, &model.SysUser{}, whereOrder)
// 	err := tx.Preload("Roles").Find(&result).Error
// 	return result, err
// }

// func (u *user) GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysUser, int64, error) {
// 	result := make([]model.SysUser, 0)
// 	tx := queryByCondition(u.db, &model.SysUser{}, whereOrder)

// 	//查询总记录数
// 	var count int64
// 	var err error
// 	err = tx.Count(&count).Error
// 	if err != nil || count == 0 {
// 		return nil, count, err
// 	}
// 	err = tx.Preload("Roles").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&result).Error
// 	return result, count, err
// }

func (u *user) Login(username, password string) (*model.SysUser, error) {
	result := model.SysUser{}
	err := u.db.Where("username = ? and password = ?", username, password).First(&result).Error
	return &result, err
}
