package v1

import "go-web/internal/auth/store"

type Service interface {
	User() UserSrv
}

type service struct {
	factory store.Factory
}

func NewService(factory store.Factory) Service {
	return &service{factory: factory}
}

func (s *service) User() UserSrv {
	return newUserService(s)
}
