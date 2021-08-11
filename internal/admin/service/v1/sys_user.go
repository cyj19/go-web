package v1

import (
	"fmt"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/cache"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/casbin/casbin/v2"
)

type SysUserSrv interface {
	Update(user *model.SysUser) error
	UpdateRoleForUser(cd *model.CreateDelete) error
	BatchDelete(ids []uint64) error
	GetByUsername(username string) (*model.SysUser, error)
	GetList(user model.SysUser) ([]model.SysUser, error)
	GetPage(userPaage model.SysUserPage) (*model.Page, error)
	Login(username, password string) (*model.SysUser, error)
}

type userService struct {
	factory  store.Factory
	enforcer *casbin.Enforcer
}

func newSysUser(srv *service) SysUserSrv {
	return &userService{
		factory:  srv.factory,
		enforcer: srv.enforcer,
	}
}

//实现SysUserSrv接口

func (u *userService) Update(param *model.SysUser) error {
	err := u.factory.SysUser().Update(param)
	if err != nil {
		return err
	}
	// 清空user相关的key
	keys := cache.Keys(param.TableName() + "*")
	cache.Del(keys...)
	return nil
}

func (u *userService) UpdateRoleForUser(cd *model.CreateDelete) error {
	// 查询记录是否存在
	srv := NewService(u.factory, u.enforcer)
	user := &model.SysUser{}
	err := srv.GetById(cd.Id, user)
	if err != nil {
		return fmt.Errorf("记录找不到：%v ", err)
	}
	err = u.factory.SysUser().UpdateRoleForUser(cd)
	if err != nil {
		return err
	}
	// 清空user相关的key
	keys := cache.Keys(user.TableName() + "*")
	cache.Del(keys...)
	return nil
}

func (u *userService) BatchDelete(ids []uint64) error {
	user := &model.SysUser{}
	err := u.factory.SysUser().BatchDelete(ids)
	if err != nil {
		return err
	}
	// 清空user相关的key
	keys := cache.Keys(user.TableName() + "*")
	cache.Del(keys...)
	return nil
}

func (u *userService) GetByUsername(username string) (*model.SysUser, error) {
	return u.factory.SysUser().GetByUsername(username)
}

func (u *userService) GetList(user model.SysUser) ([]model.SysUser, error) {
	var list []model.SysUser
	var err error
	var key string
	key = fmt.Sprintf("%s:id:%d:username:%s", user.TableName(), user.Id, user.Username)
	if user.Status != nil {
		key = fmt.Sprintf("%s:status:%t", key, *user.Status)
	}

	list = cache.GetSysUserList(key)
	if len(list) < 1 {
		whereOrders := createSysUserQueryCondition(user)
		list, err = u.factory.SysUser().GetList(whereOrders...)
		// 添加到缓存
		cache.SetSysUserList(key, list)
	}
	return list, err

}

func (u *userService) GetPage(userPage model.SysUserPage) (*model.Page, error) {
	var list []model.SysUser
	var count int64
	var err error
	var key string
	pageIndex := userPage.PageIndex
	pageSize := userPage.PageSize
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = defaultSize
	}

	// 组装key
	key = fmt.Sprintf("%s:id:%d:username:%s", userPage.TableName(), userPage.Id, userPage.Username)
	if userPage.Status != nil {
		key = fmt.Sprintf("%s:status:%t", key, *userPage.Status)
	}
	key = fmt.Sprintf("%s:pageIndex:%d:pageSize:%d", key, pageIndex, pageSize)

	list = cache.GetSysUserList(key)
	if len(list) < 1 {
		whereOrders := createSysUserQueryCondition(userPage.SysUser)
		list, err = u.factory.SysUser().GetList(whereOrders...)
		// 添加到缓存
		cache.SetSysUserList(key, list)
	}

	var userRespList []model.SysUserResponse
	util.Struct2Struct(list, &userRespList)
	page := &model.Page{
		Records:  userRespList,
		Total:    count,
		PageInfo: model.PageInfo{PageIndex: pageIndex, PageSize: pageSize},
	}
	page.SetPageNum(count)
	return page, err

}

func (u *userService) Login(username, password string) (*model.SysUser, error) {

	return u.factory.SysUser().Login(username, util.EncryptionPsw(password))

}

func createSysUserQueryCondition(param model.SysUser) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)

	if param.Id > 0 {
		v := param.Id
		whereOrders = append(whereOrders, model.WhereOrder{Where: "id = ?", Value: []interface{}{v}})
	}
	if param.Username != "" {
		v := "%" + param.Username + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "username like ?", Value: []interface{}{v}})
	}
	if param.Status != nil {
		whereOrders = append(whereOrders, model.WhereOrder{Where: "status = ?", Value: []interface{}{*param.Status}})
	}

	return whereOrders
}
