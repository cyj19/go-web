package store

import "gorm.io/gorm"

type SysApiStore interface {
}

type sysApi struct {
	db *gorm.DB
}

func newSysApi(ds *datastore) SysApiStore {
	return &sysApi{db: ds.db}
}
