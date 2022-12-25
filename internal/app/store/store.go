package store

type Store interface {
	InsertWebhook(string, interface{}) error
}
