package v1

import (
	"context"
	"fmt"
	"go-web/internal/admin/global"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/cache"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
)

type SysUserSrv interface {
	Create(ctx context.Context, values ...model.SysUser) error
	Update(ctx context.Context, user *model.SysUser) error
	UpdateRoleForUser(ctx context.Context, cd *model.CreateDelete) error
	BatchDelete(ctx context.Context, ids []uint64) error
	GetById(ctx context.Context, id uint64) (*model.SysUser, error)
	GetByUsername(ctx context.Context, username string) (*model.SysUser, error)
	GetList(ctx context.Context, user model.SysUser) ([]model.SysUser, error)
	GetPage(ctx context.Context, userPaage model.SysUserPage) (*model.Page, error)
	Login(ctx context.Context, username, password string) (*model.SysUser, error)
}

type userService struct {
	factory store.Factory
}

func newSysUser(srv *service) SysUserSrv {
	return &userService{
		factory: srv.factory,
	}
}

//实现SysUserSrv接口

func (u *userService) Create(ctx context.Context, values ...model.SysUser) error {
	err := u.factory.Create(&values)
	if err != nil {
		return err
	}
	// 清空缓存
	cleanCache(values[0].TableName() + "*")
	return nil
}

func (u *userService) Update(ctx context.Context, value *model.SysUser) error {
	err := u.factory.Update(value)
	if err != nil {
		return err
	}
	// 清空缓存
	cleanCache(value.TableName() + "*")
	return nil
}

func (u *userService) UpdateRoleForUser(ctx context.Context, cd *model.CreateDelete) error {
	// 查询记录是否存在
	user, err := u.GetById(ctx, cd.Id)
	if err != nil {
		return fmt.Errorf("记录找不到：%v ", err)
	}
	err = u.factory.SysUser().UpdateRoleForUser(cd)
	if err != nil {
		return err
	}
	// 清空缓存
	cleanCache(user.TableName() + "*")
	return nil
}

func (u *userService) BatchDelete(ctx context.Context, ids []uint64) error {
	user := new(model.SysUser)
	err := u.factory.BatchDelete(ids, user)
	if err != nil {
		return err
	}
	// 清空user相关的key
	cleanCache(user.TableName() + "*")
	return nil
}

func (u *userService) GetById(ctx context.Context, id uint64) (*model.SysUser, error) {
	value := new(model.SysUser)
	key := fmt.Sprintf("%s:id:%d", value.TableName(), id)
	err := cache.Get(global.RedisIns, key, value)
	if err != nil {
		err = u.factory.GetById(id, value)
		// 写入缓存
		cache.Set(global.RedisIns, key, value)

	}
	return value, err
}

func (u *userService) GetByUsername(ctx context.Context, username string) (*model.SysUser, error) {
	return u.factory.SysUser().GetByUsername(username)
}

func (u *userService) GetList(ctx context.Context, user model.SysUser) ([]model.SysUser, error) {
	var list []model.SysUser
	var err error
	var key string
	key = fmt.Sprintf("%s:id:%d:username:%s", user.TableName(), user.Id, user.Username)
	if user.Status != nil {
		key = fmt.Sprintf("%s:status:%t", key, *user.Status)
	}

	list = cache.GetSysUserList(global.RedisIns, key)
	if len(list) < 1 {
		whereOrders := util.GenWhereOrderByStruct(user)
		err = u.factory.GetList(&model.SysUser{}, &list, whereOrders...)
		// 添加到缓存
		cache.SetSysUserList(global.RedisIns, key, list)
	}
	return list, err

}

func (u *userService) GetPage(ctx context.Context, userPage model.SysUserPage) (*model.Page, error) {
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

	list = cache.GetSysUserList(global.RedisIns, key)
	if len(list) < 1 {
		whereOrders := util.GenWhereOrderByStruct(userPage.SysUser)
		count, err = u.factory.GetPage(pageIndex, pageSize, &model.SysUser{}, &list, whereOrders...)
		// 添加到缓存
		cache.SetSysUserList(global.RedisIns, key, list)
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

func (u *userService) Login(ctx context.Context, username, password string) (*model.SysUser, error) {

	return u.factory.SysUser().Login(username, util.EncryptionPsw(password))

}
