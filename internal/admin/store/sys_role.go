package store

import "go-web/internal/pkg/model"

type SysRoleStore interface {
	Create(r *model.SysRole) error
	Update(r *model.SysRole) error
	UpdateMenuForRole(cd *model.CreateDelete) error
	BatchDelete(ids []uint64) error
	GetByName(name string) (*model.SysRole, error)
}
