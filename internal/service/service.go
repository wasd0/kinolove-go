package service

import (
	"github.com/google/uuid"
	"kinolove/internal/service/dto"
)

type UserService interface {
	CreateUser(request dto.UserCreateRequest) (uuid.UUID, *ServErr)
	FindByUsername(username string) (dto.UserSingleResponse, *ServErr)
	Update(id uuid.UUID, request dto.UserUpdateRequest) *ServErr
}

type MovieService interface {
	CreateMovie(request dto.MovieCreateRequest) (int64, *ServErr)
	FindById(id int64) (dto.MovieSingleResponse, *ServErr)
	FindAll() (dto.MovieListResponse, *ServErr)
	Update(id int64, request dto.MovieUpdateRequest) *ServErr
}
