package domain

import (
	"time"
)

type Sensor struct {
	ID           int64     `json:"id,omitempty"`
	SoilMoisture float64   `json:"soil_moisture" binding:"required"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type SensorUsecase interface {
	FetchAll() ([]Sensor, error)
	Store(*Sensor) error
}

type SensorRepository interface {
	FetchAll() ([]Sensor, error)
	Store(*Sensor) error
}
