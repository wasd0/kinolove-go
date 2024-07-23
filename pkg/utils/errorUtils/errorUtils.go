package errorUtils

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func GetPgxErr(err error, op string, defaultMsg string) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return fmt.Errorf("%s >  %s", op, pgErr.Message)
	}

	return fmt.Errorf("%s > %s", op, defaultMsg)
}
