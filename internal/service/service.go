package service

import (
	"github.com/google/uuid"
	"kinolove/internal/service/dto"
)

type UserService interface {
	CreateUser(request dto.UserCreateRequest) (uuid.UUID, error)
	FindByUsername(username string) (dto.UserSingleResponse, error)
	Update(id uuid.UUID, request dto.UserUpdateRequest) error
}

type MovieService interface {
	CreateMovie(request dto.MovieCreateRequest) (int64, error)
	FindById(id int64) (dto.MovieSingleResponse, error)
	FindAll() (dto.MovieListResponse, error)
	Update(id int64, request dto.MovieUpdateRequest) error
}
