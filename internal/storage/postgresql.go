package storage

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"kinolove/pkg/constants"
	"kinolove/pkg/logger"
	"kinolove/pkg/utils/app"
	"log"
	"os"
)

const DbDriver = "pgx"

type PgStorage struct {
	Db *sql.DB
}

func MustOpenPostgres() (*PgStorage, app.Callback) {
	dbUrl := os.Getenv(constants.EnvDbUrl)

	db, err := sql.Open(DbDriver, dbUrl)

	if err != nil {
		logger.Log().Fatalf(err, "%s Failed to open connection to database", constants.OpenConnect)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err, "postgres ping failed")
	}

	return &PgStorage{Db: db}, func(ctx context.Context) error {
		logger.Log().Info("Database closing...")
		return db.Close()
	}
}
