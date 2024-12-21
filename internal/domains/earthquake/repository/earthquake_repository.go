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

func (r *EarthquakeRepository) FindMoreThanMagnitude(earthquakes *[]entity.Earthquake, value float64) error {
	return r.DB.Where("magnitude >= ?", value).Find(&earthquakes).Error
}

func (r *EarthquakeRepository) FindLessThanMagnitude(earthquakes *[]entity.Earthquake, value float64) error {
	return r.DB.Where("magnitude <= ?", value).Find(&earthquakes).Error
}

func (r *EarthquakeRepository) FindMoreThanDepth(earthquakes *[]entity.Earthquake, value int64) error {
	return r.DB.Where("depth >= ?", value).Find(&earthquakes).Error
}

func (r *EarthquakeRepository) FindLessThanDepth(earthquakes *[]entity.Earthquake, value int64) error {
	return r.DB.Where("depth <= ?", value).Find(&earthquakes).Error
}
