package store

type Factory interface {
	User() UserStore
}
