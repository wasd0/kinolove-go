package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	. "kinolove/internal/entity/.gen/kinolove/public/table"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/errorUtils"
)

type GenreRepositoryImpl struct {
	db *sql.DB
}

func NewGenreRepository(sqlDb *sql.DB) *GenreRepositoryImpl {
	return &GenreRepositoryImpl{db: sqlDb}
}

func (g *GenreRepositoryImpl) Save(entity *model.Genres) error {
	stmt := Genres.INSERT(Genres.MutableColumns).MODEL(entity).RETURNING(Genres.AllColumns)
	err := stmt.Query(g.db, entity)

	if err != nil {
		err = errorUtils.GetPgxErr(err, constants.Insert, "Failed save entity")
	}

	return nil
}

func (g *GenreRepositoryImpl) FindAllByMovieId(IdTitle int64) ([]model.Genres, error) {
	genres := make([]model.Genres, 0)
	stmt :=
		SELECT(Genres.AllColumns).
			FROM(Genres.
				INNER_JOIN(GenresTitles, Genres.ID.EQ(GenresTitles.GenreID).AND(GenresTitles.TitleID.EQ(Int64(IdTitle)))))

	if err := stmt.Query(g.db, &genres); err != nil {
		return nil, err
	}

	return genres, nil
}
