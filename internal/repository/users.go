package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"kinolove/internal/common/constants"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/entity/.gen/kinolove/public/table"
)

type UserPgRepo[ENTITY model.Users, ID uuid.UUID] struct {
	db    *sql.DB
	users *table.UsersTable
}

func NewUserRepository(sqlDb *sql.DB) *UserPgRepo[model.Users, uuid.UUID] {
	return &UserPgRepo[model.Users, uuid.UUID]{db: sqlDb, users: table.Users}
}

func (r *UserPgRepo[ENTITY, ID]) GetById(id ID) (*ENTITY, error) {
	var usr ENTITY

	stmt := postgres.
		SELECT(r.users.AllColumns).
		FROM(r.users).
		WHERE(r.users.ID.EQ(postgres.UUID(uuid.UUID(id))))

	err := stmt.Query(r.db, &usr)

	if err != nil {
		err = fmt.Errorf("%s > user not found by id %s", constants.Select, id)
	}

	return &usr, err

}

func (r *UserPgRepo[ENTITY, ID]) GetByUsername(username string) (*ENTITY, error) {
	var usr ENTITY

	stmt := postgres.
		SELECT(r.users.AllColumns).
		FROM(r.users).
		WHERE(r.users.Username.EQ(postgres.String(username)))

	err := stmt.Query(r.db, &usr)

	if err != nil {
		err = fmt.Errorf("%s > user not found by username %s", constants.Select, username)
	}

	return &usr, err
}

func (r *UserPgRepo[ENTITY, ID]) Save(entity *ENTITY) error {
	stmt := r.users.INSERT(r.users.Username, r.users.Password).MODEL(entity).RETURNING(r.users.AllColumns)
	err := stmt.Query(r.db, entity)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("%s >  %s", constants.Insert, pgErr.Message)
		}

		return fmt.Errorf("%s > %s", constants.Insert, err)
	}

	return nil
}
