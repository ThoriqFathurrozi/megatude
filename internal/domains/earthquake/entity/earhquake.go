package entity

import (
	"github.com/ThoriqFathurrozi/megatude/internal/types"
)

// Coordinates struct
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Earthquake struct {
	types.Entity
	Datetime    string      `json:"datetime"`
	Magnitude   float64     `json:"magnitude"`
	Depth       float64     `json:"depth"`
	Coordinates Coordinates `json:"coordinates"`
	Longitude   string      `json:"longitude"`
	Latitude    string      `json:"latitude"`
	Location    string      `json:"location"`
}
