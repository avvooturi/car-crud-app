package storage

import "github.com/avvooturi/car-crud-app/pkg/models"

type DB interface {
	FetchAll() []models.Car
	FetchOne(id string) (models.Car, error)
	Create(id string, cars ...models.Car) ([]models.Car, error)
	Delete(id string) error
	Update(id string, car models.Car) (models.Car, error)
}