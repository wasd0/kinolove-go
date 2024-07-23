package repository

import (
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/entity/.gen/kinolove/public/table"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/errorUtils"
)

type MoviePgRepo struct {
	db     *sql.DB
	movies *table.MoviesTable
}

func NewMoviesRepository(sqlDb *sql.DB) *MoviePgRepo {
	return &MoviePgRepo{db: sqlDb, movies: table.Movies}
}

func (m *MoviePgRepo) GetById(id int64) (*model.Movies, error) {
	var movie model.Movies

	stmt := postgres.
		SELECT(m.movies.AllColumns).
		FROM(m.movies).
		WHERE(m.movies.ID.EQ(postgres.Int64(id)))

	err := stmt.Query(m.db, &movie)

	if err != nil {
		return nil, errorUtils.GetPgxErr(err, constants.Select, "error while fetching movies")
	}

	return &movie, nil
}

func (m *MoviePgRepo) Save(entity *model.Movies) error {
	stmt := m.movies.INSERT(m.movies.MutableColumns).MODEL(entity).RETURNING(m.movies.AllColumns)
	err := stmt.Query(m.db, entity)

	if err != nil {
		return errorUtils.GetPgxErr(err, constants.Insert, "error while save movie")
	}

	return nil
}

func (m *MoviePgRepo) Update(entity *model.Movies) error {
	stmt := m.movies.UPDATE(m.movies.MutableColumns).
		MODEL(entity).
		WHERE(m.movies.ID.EQ(postgres.Int64(entity.ID)))

	_, err := stmt.Exec(m.db)

	if err != nil {
		return errorUtils.GetPgxErr(err, constants.Update, "error while updating movies")
	}

	return nil
}

func (m *MoviePgRepo) FindAll() (*[]*model.Movies, error) {
	movies := make([]*model.Movies, 0)

	stmt := postgres.SELECT(m.movies.AllColumns).FROM(m.movies)
	err := stmt.Query(m.db, &movies)

	if err != nil {
		return nil, errorUtils.GetPgxErr(err, constants.Select, "Failed find all movies")
	}

	return &movies, nil
}
