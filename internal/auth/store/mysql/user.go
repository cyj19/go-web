package mysql

import (
	"go-web/internal/auth/store"
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func newUser(d *database) store.UserStore {
	return &user{db: d.db}
}

func (u *user) Login(username, password string) (*model.SysUser, error) {
	user := &model.SysUser{}
	err := u.db.Where("username = ? and password = ?", username, password).First(user).Error
	return user, err
}
