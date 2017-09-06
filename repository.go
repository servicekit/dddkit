package dddkit

import (
	"github.com/servicekit/dddkit/storage"
)

type Repository interface {
	GetByID(id string) (Entity, error)
	Get(query *storage.Query) ([]Entity, error)
	Insert(e Entity) error
	Update(e Entity) error
	Delete(e Entity) error
}
