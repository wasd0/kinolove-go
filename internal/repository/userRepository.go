package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	. "kinolove/internal/entity/.gen/kinolove/public/table"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/errorUtils"
)

type UserPgRepo struct {
	db *sql.DB
}

func NewUserRepository(sqlDb *sql.DB) *UserPgRepo {
	return &UserPgRepo{db: sqlDb}
}

func (r *UserPgRepo) GetById(id uuid.UUID) (*model.Users, error) {
	var usr model.Users

	stmt := SELECT(Users.AllColumns).
		FROM(Users).
		WHERE(Users.ID.EQ(UUID(id)))

	err := stmt.Query(r.db, &usr)

	if err != nil {
		return nil, err
	}

	return &usr, err

}

func (r *UserPgRepo) GetByUsername(username string) (*model.Users, error) {
	var usr model.Users

	stmt :=
		SELECT(Users.AllColumns).
			FROM(Users).
			WHERE(Users.Username.EQ(String(username)))

	err := stmt.Query(r.db, &usr)

	if err != nil {
		return nil, err
	}

	return &usr, err
}

func (r *UserPgRepo) Save(entity *model.Users) error {
	stmt :=
		Users.INSERT(Users.Username, Users.Password).
			MODEL(entity).
			RETURNING(Users.AllColumns)
	err := stmt.Query(r.db, entity)
	if err != nil {
		err = errorUtils.GetPgxErr(err, constants.Insert, "Failed save entity")
	}

	return err
}

func (r *UserPgRepo) ExistsByUsername(username string) (bool, error) {
	stmt := SELECT(COUNT(Users.ID).GT(Int(0)).
		AS("Exists")).
		FROM(Users).
		WHERE(Users.Username.EQ(String(username)))

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
	stmt := Users.UPDATE(Users.MutableColumns).
		MODEL(entity).
		WHERE(Users.ID.EQ(UUID(entity.ID)))

	_, err := stmt.Exec(r.db)

	if err != nil {
		return errorUtils.GetPgxErr(err, constants.Update, "error while updating movies")
	}

	return nil
}

func (r *UserPgRepo) FindAll() (*[]*model.Users, error) {
	users := make([]*model.Users, 0)

	stmt := SELECT(Users.AllColumns).FROM(Users)
	err := stmt.Query(r.db, &users)

	if err != nil {
		return nil, err
	}

	return &users, nil
}
