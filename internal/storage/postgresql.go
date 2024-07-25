package storage

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"kinolove/pkg/constants"
	"kinolove/pkg/logger"
	"kinolove/pkg/utils/app"
	"os"
)

const DbDriver = "pgx"

type PgStorage struct {
	Db *sql.DB
}

func MustOpenPostgres(log logger.Common) (*PgStorage, app.Callback) {
	dbUrl := os.Getenv(constants.EnvDbUrl)

	db, err := sql.Open(DbDriver, dbUrl)

	if err != nil {
		log.Fatalf(err, "%s Failed to open connection to database", constants.OpenConnect)
	}

	return &PgStorage{Db: db}, func(ctx context.Context) error {
		log.Info("Database closing...")
		return db.Close()
	}
}
