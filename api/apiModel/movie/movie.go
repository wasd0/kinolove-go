package movie

import (
	"github.com/pkg/errors"
	"kinolove/internal/service/dto"
	"net/http"
)

type ReqMovieCreate struct {
	dto.MovieCreateRequest
}

func (m ReqMovieCreate) Bind(_ *http.Request) error {
	if len(m.Title) == 0 {
		return errors.New("Invalid title")
	}

	return nil
}

type ResMovieFindAll struct {
	Data dto.MovieListResponse `json:"data"`
}

func (m *ResMovieFindAll) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type ReqMovieUpdate struct {
	dto.MovieUpdateRequest
}

func (r *ReqMovieUpdate) Bind(_ *http.Request) error {
	return nil
}
