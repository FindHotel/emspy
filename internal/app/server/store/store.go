package store

type Store interface {
	Insert(interface{}) error
}
