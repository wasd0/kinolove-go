package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/entity/.gen/kinolove/public/table"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/errorUtils"
)

type UserPgRepo struct {
	db    *sql.DB
	users *table.UsersTable
}

func NewUserRepository(sqlDb *sql.DB) *UserPgRepo {
	return &UserPgRepo{db: sqlDb, users: table.Users}
}

func (r *UserPgRepo) GetById(id uuid.UUID) (*model.Users, error) {
	var usr model.Users

	stmt := postgres.
		SELECT(r.users.AllColumns).
		FROM(r.users).
		WHERE(r.users.ID.EQ(postgres.UUID(id)))

	err := stmt.Query(r.db, &usr)

	if err != nil {
		err = fmt.Errorf("%s > user not found by id %s", constants.Select, id)
	}

	return &usr, err

}

func (r *UserPgRepo) GetByUsername(username string) (*model.Users, error) {
	var usr model.Users

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

func (r *UserPgRepo) Save(entity *model.Users) error {
	stmt := r.users.
		INSERT(r.users.Username, r.users.Password).
		MODEL(entity).
		RETURNING(r.users.AllColumns)
	err := stmt.Query(r.db, entity)
	if err != nil {
		err = errorUtils.GetPgxErr(err, constants.Insert, "Failed save entity")
	}

	return err
}

func (r *UserPgRepo) ExistsByUsername(username string) (bool, error) {
	stmt := postgres.
		SELECT(postgres.COUNT(r.users.ID).GT(postgres.Int(0)).
			AS("Exists")).
		FROM(r.users).
		WHERE(r.users.Username.EQ(postgres.String(username)))

	var res struct {
		Exists bool
	}
	err := stmt.Query(r.db, &res)

	if err != nil {
		err = errorUtils.GetPgxErr(err, constants.Select, "Failed check user exists by username")
	}

	return res.Exists, err
}

func (r *UserPgRepo) Update(entity *model.Users) error {
	stmt := r.users.UPDATE(r.users.MutableColumns).
		MODEL(entity).
		WHERE(r.users.ID.EQ(postgres.UUID(entity.ID)))

	_, err := stmt.Exec(r.db)

	if err != nil {
		return errorUtils.GetPgxErr(err, constants.Update, "error while updating movies")
	}

	return nil
}

func (r *UserPgRepo) FindAll() (*[]*model.Users, error) {
	users := make([]*model.Users, 0)

	stmt := postgres.SELECT(r.users.AllColumns).FROM(r.users)
	err := stmt.Query(r.db, &users)

	if err != nil {
		return nil, errorUtils.GetPgxErr(err, constants.Select, "Failed find all users")
	}

	return &users, nil
}
