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
func (u *user) Create(user *model.SysUser) error {
	return u.db.Create(user).Error
}

func (u *user) Update(user *model.SysUser) error {
	return u.db.Save(user).Error
}

func (u *user) Delete(id uint64) error {
	return delete(u.db, id, &model.SysUser{})
}

func (u *user) DeleteBatch(ids []uint64) error {
	return deleteBatch(u.db, ids, &model.SysUser{})
}

func (u *user) GetByUsername(username string) (*model.SysUser, error) {
	result := &model.SysUser{}
	err := u.db.Where("username = ?", username).First(result).Error
	return result, err
}

func (u *user) List(whereOrder ...model.WhereOrder) ([]model.SysUser, error) {
	result := make([]model.SysUser, 0)
	tx := queryByCondition(u.db, &model.SysUser{}, whereOrder)
	err := tx.Find(&result).Error
	return result, err
}

func (u *user) GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysUser, int64, error) {
	result := make([]model.SysUser, 0)
	tx := queryByCondition(u.db, &model.SysUser{}, whereOrder)

	//查询总记录数
	var count int64
	var err error
	err = tx.Count(&count).Error
	if err != nil || count == 0 {
		return nil, count, err
	}
	err = tx.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&result).Error
	return result, count, err
}

func (u *user) Login(username, password string) (*model.SysUser, error) {
	result := model.SysUser{}
	err := u.db.Where("username = ? and password = ?", username, password).First(&result).Error
	return &result, err
}
