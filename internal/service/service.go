package service

import (
	"github.com/google/uuid"
	"kinolove/internal/service/dto/user"
)

type UserService interface {
	CreateUser(request user.CreateRequest) (uuid.UUID, error)
	FindByUsername(username string) (user.SingleResponse, error)
}
