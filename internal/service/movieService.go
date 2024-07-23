package service

import (
	"fmt"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/repository"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/mapper"
)

type MovieServiceImpl struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieServiceImpl {
	return &MovieServiceImpl{movieRepo: &repo}
}

func (m *MovieServiceImpl) CreateMovie(request dto.MovieCreateRequest) (int64, error) {
	repo := *m.movieRepo

	if len(request.Title) == 0 {
		return -1, fmt.Errorf("title can not be empty")
	}

	movie := model.Movies{
		Title: request.Title,
	}

	err := repo.Save(&movie)
	if err != nil {
		return 0, err
	}

	return movie.ID, nil
}

func (m *MovieServiceImpl) FindById(id int64) (dto.MovieSingleResponse, error) {
	repo := *m.movieRepo

	movie, err := repo.GetById(id)
	if err != nil {
		return dto.MovieSingleResponse{}, err
	}

	return mapper.MapMovieToSingleResponse(movie), nil
}

func (m *MovieServiceImpl) FindAll() (dto.MovieListResponse, error) {
	repo := *m.movieRepo
	movies, err := repo.FindAll()

	if err != nil {
		return dto.MovieListResponse{}, err
	}

	data := make([]dto.MovieItemData, 0, len(*movies))

	for _, movie := range *movies {
		data = append(data, mapper.MapMovieToItemData(movie))
	}

	return dto.MovieListResponse{Data: data}, nil
}

func (m *MovieServiceImpl) Update(id int64, request dto.MovieUpdateRequest) error {
	repo := *m.movieRepo

	movie, err := repo.GetById(id)

	if err != nil {
		return err
	}

	err = mapper.MapUpdateRequestToMovie(&request, movie)

	if err != nil {
		return err
	}

	return repo.Update(movie)
}
