package entity

import (
	"time"

	"github.com/ThoriqFathurrozi/megatude/internal/types"
)

type Earthquake struct {
	types.Entity
	Datetime    time.Time `gorm:"unique" json:"datetime"`
	Magnitude   float64   `json:"magnitude"`
	Depth       int64     `json:"depth"`
	Coordinates string    `json:"coordinates_id"`
	Longitude   string    `json:"longitude"`
	Latitude    string    `json:"latitude"`
	Location    string    `json:"location"`
}

func NewEarthquake() *Earthquake {
	return &Earthquake{}
}

func (e Earthquake) TableName() string {
	return "earthquake"
}
