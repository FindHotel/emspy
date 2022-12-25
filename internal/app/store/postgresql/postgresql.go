package postgresql

import (
	"database/sql"

	"github.com/FindHotel/emspy/internal/app/store"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PostgresqlStore struct {
	db *bun.DB
}

func New(dsn string) (store.Store, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return &PostgresqlStore{db: db}, nil
}

func (s *PostgresqlStore) InsertWebhook(source string, record interface{}) error {

	return nil
}
