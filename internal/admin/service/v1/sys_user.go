package v1

import (
	"fmt"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
)

type SysUserSrv interface {
	Update(user *model.SysUser) error
	UpdateRoleForUser(cd *model.CreateDelete) error
	BatchDelete(ids []uint64) error
	GetById(id uint64) (*model.SysUser, error)
	GetByUsername(username string) (*model.SysUser, error)
	GetList(whereOrders ...model.WhereOrder) ([]model.SysUser, error)
	GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) ([]model.SysUser, int64, error)
	Login(username, password string) (*model.SysUser, error)
}

type userService struct {
	factory store.Factory
}

func newSysUser(srv *service) SysUserSrv {
	return &userService{factory: srv.factory}
}

//实现SysUserSrv接口

func (u *userService) Update(param *model.SysUser) error {
	return u.factory.SysUser().Update(param)
}

func (u *userService) UpdateRoleForUser(cd *model.CreateDelete) error {
	// 查询记录是否存在
	_, err := u.GetById(cd.Id)
	if err != nil {
		return fmt.Errorf("记录找不到：%v ", err)
	}
	return u.factory.SysUser().UpdateRoleForUser(cd)
}

func (u *userService) BatchDelete(ids []uint64) error {
	return u.factory.SysUser().BatchDelete(ids)
}

func (u *userService) GetById(id uint64) (*model.SysUser, error) {
	return u.factory.SysUser().GetById(id)
}

func (u *userService) GetByUsername(username string) (*model.SysUser, error) {
	return u.factory.SysUser().GetByUsername(username)
}

func (u *userService) GetList(whereOrders ...model.WhereOrder) ([]model.SysUser, error) {
	//构建查询条件
	//whereOrder := createSysUserQueryCondition(user)
	return u.factory.SysUser().GetList(whereOrders...)

}

func (u *userService) GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) ([]model.SysUser, int64, error) {
	return u.factory.SysUser().GetPage(pageIndex, pageSize, whereOrders...)

}

func (u *userService) Login(username, password string) (*model.SysUser, error) {
	return u.factory.SysUser().Login(username, password)

}
