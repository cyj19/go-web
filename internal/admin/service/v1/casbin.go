package v1

import "go-web/internal/pkg/model"

type CasbinSrv interface {
	Create(sc *model.SysCasbin) error
}

// casbin维护了用户-角色-权限
// 缺少维护用户-角色
