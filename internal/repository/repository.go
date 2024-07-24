package repository

import (
	"github.com/google/uuid"
	"kinolove/internal/entity/.gen/kinolove/public/model"
)

type Repository[ENTITY any, ID any] interface {
	GetById(id ID) (*ENTITY, error)
	Save(entity *ENTITY) error
	Update(entity *ENTITY) error
	FindAll() (*[]*ENTITY, error)
}

type UserRepository interface {
	Repository[model.Users, uuid.UUID]
	GetByUsername(username string) (*model.Users, error)
	ExistsByUsername(username string) (bool, error)
}

type MovieRepository interface {
	Repository[model.Movies, int64]
}
