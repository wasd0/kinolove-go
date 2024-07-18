package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx5"
	"kinolove/internal/common/constants"
	"kinolove/internal/logger"
	"kinolove/internal/storage"
	"os"
)

type PgStorage struct {
	db  *ksql.DB
	log logger.Common
}

func (p PgStorage) SelectOne(entityPtr interface{}, query string, params ...interface{}) error {
	return p.db.QueryOne(context.Background(), entityPtr, query, params...)
}

func (p PgStorage) SelectAll(entitiesPtr interface{}, query string, params ...interface{}) error {
	return p.db.Query(context.Background(), entitiesPtr, query, params...)
}

func (p PgStorage) Insert(table string, entityPtr interface{}) error {
	ksqlTable := ksql.NewTable(table)
	return p.db.Insert(context.Background(), ksqlTable, entityPtr)
}

func (p PgStorage) Update(table string, entityPtr interface{}) error {
	ksqlTable := ksql.NewTable(table)
	return p.db.Patch(context.Background(), ksqlTable, entityPtr)
}

func (p PgStorage) Delete(table string, idOrEntityPtr interface{}) error {
	ksqlTable := ksql.NewTable(table)
	return p.db.Delete(context.Background(), ksqlTable, idOrEntityPtr)
}

func (p PgStorage) Transaction(fn func(storage storage.Storage) error) error {
	return p.db.Transaction(context.Background(), func(provider ksql.Provider) error {
		//todo implement
		return nil
	})
}

func MustOpenConnect(log logger.Common) (storage.Storage, func()) {
	op := constants.OpenConnect

	pool, err := pgxpool.New(context.Background(), os.Getenv(constants.EnvDbUrl))

	if err != nil {
		log.Fatalf(err, "[%s] failed to connect to database", op)
	}

	db, err := kpgx.NewFromPgxPool(pool)

	if err != nil {
		log.Fatalf(err, "[%s] failed to use sql dialect", op)
	}

	log.Info("Postgresql database connected")

	return &PgStorage{db: &db, log: log}, func() {
		operation := constants.CloseConnect
		closeErr := db.Close()
		if closeErr != nil {
			log.Fatalf(err, "[%s] failed to close database connection", operation)
		}
	}
}
