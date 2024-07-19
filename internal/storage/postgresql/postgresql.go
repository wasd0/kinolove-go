package postgresql

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"kinolove/internal/common/constants"
	"kinolove/internal/logger"
	"os"
)

const DbDriver = "pgx"

type PgStorage struct {
	Db *sql.DB
}

func MustOpenConnect(log logger.Common) (*PgStorage, func()) {
	dbUrl := os.Getenv(constants.EnvDbUrl)

	db, err := sql.Open(DbDriver, dbUrl)

	if err != nil {
		log.Fatalf(err, "%s Failed to open connection to database", constants.OpenConnect)
	}

	return &PgStorage{Db: db}, func() {
		err = db.Close()

		if err != nil {
			log.Fatalf(err, "%s failed to close database", constants.CloseConnect)
		}
	}
}
