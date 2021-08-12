package store

import "go-web/internal/pkg/model"

type SysRoleStore interface {
	UpdateMenuForRole(cd *model.CreateDelete) error
	GetByName(name string) (*model.SysRole, error)
}
