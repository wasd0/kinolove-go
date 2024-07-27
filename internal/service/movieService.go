package service

import (
	"fmt"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/pkg/errors"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/repository"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/mapper"
)

type MovieServiceImpl struct {
	movieRepo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieServiceImpl {
	return &MovieServiceImpl{movieRepo: repo}
}

func (m *MovieServiceImpl) CreateMovie(request dto.MovieCreateRequest) (int64, *ServErr) {
	if len(request.Title) == 0 {
		return -1, BadRequest(errors.New("Wring title"), "movie title can not be empty")
	}

	movie := model.Movies{
		Title: request.Title,
	}

	err := m.movieRepo.Save(&movie)
	if err != nil {
		return 0, InternalError(err)
	}

	return movie.ID, nil
}

func (m *MovieServiceImpl) FindById(id int64) (dto.MovieSingleResponse, *ServErr) {
	movie, err := m.movieRepo.GetById(id)
	if err != nil {
		msg := fmt.Sprintf("Movie with id %d not found", id)
		return dto.MovieSingleResponse{}, BadRequest(err, msg)
	}

	return mapper.MapMovieToSingleResponse(movie), nil
}

func (m *MovieServiceImpl) FindAll() (dto.MovieListResponse, *ServErr) {
	movies, err := m.movieRepo.FindAll()

	if err != nil && errors.Is(err, qrm.ErrNoRows) {
		return dto.MovieListResponse{}, InternalError(err)
	}

	data := make([]dto.MovieItemData, 0, len(*movies))

	for _, movie := range *movies {
		data = append(data, mapper.MapMovieToItemData(movie))
	}

	return dto.MovieListResponse{Movies: data}, nil
}

func (m *MovieServiceImpl) Update(id int64, request dto.MovieUpdateRequest) *ServErr {
	movie, err := m.movieRepo.GetById(id)

	if err != nil {
		msg := fmt.Sprintf("Movie with id %d not found", id)
		return BadRequest(err, msg)
	}

	err = mapper.MapUpdateRequestToMovie(&request, movie)

	if err != nil {
		return InternalError(err)
	}

	if err = m.movieRepo.Update(movie); err != nil {
		return InternalError(err)
	}

	return nil
}
