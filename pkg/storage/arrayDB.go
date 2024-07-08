package storage

import (
	"fmt"

	"github.com/avvooturi/car-crud-app/pkg/models"
)

type arrDB struct {
	Data map[string]models.Car
}

func (db arrDB) FetchAll() []models.Car {
	return []models.Car{}
}

func (db arrDB) FetchOne(id string) (models.Car, error) {
	v, found := db.Data[id]
	if !found {
		return models.Car{}, fmt.Errorf("id dne")
	}
	return v, nil
}

func (db *arrDB) Create(id string, cars ...models.Car) ([]models.Car, error) {
	return []models.Car{}, nil
}

func (db *arrDB) Delete(id string) error {
	return nil
}

func (db *arrDB) Update(id string, car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func NewArrayDB() DB {
	return &arrDB{Data: make(map[string]models.Car)}
}
