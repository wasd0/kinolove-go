package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	. "kinolove/internal/entity/.gen/kinolove/public/table"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/errorUtils"
)

type MoviePgRepo struct {
	db *sql.DB
}

func NewMoviesRepository(sqlDb *sql.DB) *MoviePgRepo {
	return &MoviePgRepo{db: sqlDb}
}

func (m *MoviePgRepo) GetById(id int64) (*model.Movies, error) {
	var movie model.Movies

	stmt := SELECT(Movies.AllColumns).
		FROM(Movies).
		WHERE(Movies.ID.EQ(Int64(id)))

	err := stmt.Query(m.db, &movie)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *MoviePgRepo) Save(entity *model.Movies) error {
	stmt := Movies.INSERT(Movies.MutableColumns).MODEL(entity).RETURNING(Movies.AllColumns)
	err := stmt.Query(m.db, entity)

	if err != nil {
		return errorUtils.GetPgxErr(err, constants.Insert, "error while save movie")
	}

	return nil
}

func (m *MoviePgRepo) Update(entity *model.Movies) error {
	stmt := Movies.UPDATE(Movies.MutableColumns).
		MODEL(entity).
		WHERE(Movies.ID.EQ(Int64(entity.ID)))

	_, err := stmt.Exec(m.db)

	if err != nil {
		return errorUtils.GetPgxErr(err, constants.Update, "error while updating movies")
	}

	return nil
}

func (m *MoviePgRepo) FindAll() (*[]*model.Movies, error) {
	movies := make([]*model.Movies, 0)

	stmt := SELECT(Movies.AllColumns).FROM(Movies)
	err := stmt.Query(m.db, &movies)

	if err != nil {
		return nil, err
	}

	return &movies, nil
}
