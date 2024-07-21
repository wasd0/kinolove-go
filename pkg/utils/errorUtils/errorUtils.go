package errorUtils

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"kinolove/pkg/constants"
)

func TryGetPgxErr(err error, defaultMsg string) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return fmt.Errorf("%s >  %s", constants.Insert, pgErr.Message)
	}

	return fmt.Errorf("%s > %s", constants.Insert, defaultMsg)
}
