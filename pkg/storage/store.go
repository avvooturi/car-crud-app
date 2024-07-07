package storage

import "car-api/pkg/models"

type DB interface {
	FetchAll() []models.Car
	FetchOne(id string) (models.Car, error)
	Create(id string, cars ...models.Car) ([]models.Car, error)
	Delete(id string) error
	Update(id string, car models.Car) (models.Car, error)
}

type arrDB struct {
	Data map[string] models.Car
}
