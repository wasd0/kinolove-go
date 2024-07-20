package repository

type Repository[ENTITY any, ID any] interface {
	GetById(id ID) (*ENTITY, error)
	Save(entity *ENTITY) error
}
