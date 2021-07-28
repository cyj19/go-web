package v1

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
)

type SysUserSrv interface {
	Create(user *model.SysUser) error
	Update(user *model.SysUser) error
	Delete(id uint64) error
	DeleteBatch(ids []uint64) error
	GetByUsername(username string) (*model.SysUser, error)
	List(user *model.SysUser) ([]model.SysUser, error)
	GetPage(userPage *model.SysUserPage) ([]model.SysUser, int64, error)
	Login(username, password string) (*model.SysUser, error)
}

type userService struct {
	factory store.Factory
}

func newSysUser(srv *service) SysUserSrv {
	return &userService{factory: srv.factory}
}

//实现SysUserSrv接口

func (u *userService) Create(user *model.SysUser) error {
	return u.factory.SysUser().Create(user)
}

func (u *userService) Delete(id uint64) error {
	return u.factory.SysUser().Delete(id)

}

func (u *userService) DeleteBatch(ids []uint64) error {
	return u.factory.SysUser().DeleteBatch(ids)
}

func (u *userService) Update(param *model.SysUser) error {
	return u.factory.SysUser().Update(param)
}

func (u *userService) GetByUsername(username string) (*model.SysUser, error) {
	return u.factory.SysUser().GetByUsername(username)
}

func (u *userService) List(user *model.SysUser) ([]model.SysUser, error) {
	//构建查询条件
	whereOrder := createSysUserCondition(user)
	return u.factory.SysUser().List(whereOrder...)

}

func (u *userService) GetPage(userPage *model.SysUserPage) ([]model.SysUser, int64, error) {
	//构建查询条件
	whereOrder := createSysUserCondition(&userPage.SysUser)
	pageIndex := userPage.PageIndex
	pageSize := userPage.PageSize
	if pageIndex < 1 {
		pageIndex = 1
	}
	return u.factory.SysUser().GetPage(pageIndex, pageSize, whereOrder...)

}

func (u *userService) Login(username, password string) (*model.SysUser, error) {
	return u.factory.SysUser().Login(username, password)

}

func createSysUserCondition(param *model.SysUser) []model.WhereOrder {
	var whereOrder []model.WhereOrder
	if param != nil {
		if param.Username != "" {
			v := "%" + param.Username + "%"
			whereOrder = append(whereOrder, model.WhereOrder{Where: "username like ?", Value: []interface{}{v}})
		}
		if param.Status != nil {
			whereOrder = append(whereOrder, model.WhereOrder{Where: "status = ?", Value: []interface{}{param.Status}})
		}
	}

	return whereOrder
}
