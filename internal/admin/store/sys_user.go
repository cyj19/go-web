package store

import (
	"github.com/vagaryer/go-web/internal/pkg/model"

	"gorm.io/gorm"
)

//SysUserStore defines the user storage interface.
type SysUserStore interface {
	UpdateRoleForUser(cd *model.CreateDelete) error
	GetByUsername(username string) (*model.SysUser, error)
	Login(username, password string) (*model.SysUser, error)
}

type user struct {
	db *gorm.DB
}

func newSysUser(ds *datastore) SysUserStore {
	return &user{db: ds.db}
}

var _ SysUserStore = (*user)(nil)

//实现SysUserStore接口

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

func (u *user) GetByUsername(username string) (*model.SysUser, error) {
	result := model.SysUser{}
	err := u.db.Preload("Roles").Where("username = ?", username).First(&result).Error
	return &result, err
}

func (u *user) Login(username, password string) (*model.SysUser, error) {
	result := model.SysUser{}
	err := u.db.Where("username = ? and password = ?", username, password).First(&result).Error
	return &result, err
}
