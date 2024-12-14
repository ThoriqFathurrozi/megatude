package repository

import (
	"github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake/entity"
	"github.com/ThoriqFathurrozi/megatude/internal/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EarthquakeRepository struct {
	types.DBRepository[entity.Earthquake]
}

func NewEarthquakeRepository(db *gorm.DB) *EarthquakeRepository {
	return &EarthquakeRepository{
		DBRepository: types.DBRepository[entity.Earthquake]{
			DB: db,
		},
	}
}

func (r *EarthquakeRepository) Create(earthquake *entity.Earthquake) error {
	return r.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(earthquake).Error
}

func (r *EarthquakeRepository) FindAll(earthquakes *[]entity.Earthquake) error {
	return r.DB.Find(earthquakes).Error
}
