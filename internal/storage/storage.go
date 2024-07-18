package storage

type Storage interface {
	SelectOne(entityPtr interface{}, query string, params ...interface{}) error
	SelectAll(entitiesPtr interface{}, query string, params ...interface{}) error
	Insert(table string, entityPtr interface{}) error
	Update(table string, entityPtr interface{}) error
	Delete(table string, idOrEntityPtr interface{}) error
	Transaction(fn func(storage Storage) error) error
}
