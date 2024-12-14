package types

import (
	"gorm.io/gorm"
)

type DBRepository[T any] struct {
	DB *gorm.DB
}

func (r *DBRepository[T]) Create(entity T) error {
	return r.DB.Create(entity).Error
}

func (r *DBRepository[T]) FindAll(entities *[]T) error {
	return r.DB.Find(entities).Order("created_at desc").Error
}
