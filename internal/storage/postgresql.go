package storage

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"kinolove/pkg/constants"
	"kinolove/pkg/logger"
	"os"
)

const DbDriver = "pgx"

type PgStorage struct {
	Db *sql.DB
}

func MustOpenPostgres(log logger.Common) (*PgStorage, func()) {
	dbUrl := os.Getenv(constants.EnvDbUrl)

	db, err := sql.Open(DbDriver, dbUrl)

	if err != nil {
		log.Fatalf(err, "%s Failed to open connection to database", constants.OpenConnect)
	}

	return &PgStorage{Db: db}, func() {
		closeDb(log, db)
	}
}

func closeDb(log logger.Common, db *sql.DB) {
	err := db.Close()

	if err != nil {
		log.Fatalf(err, "%s failed to close database", constants.CloseConnect)
	}

}
