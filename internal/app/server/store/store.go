package store

type Store interface {
	InsertWebhook(interface{}) error
}
