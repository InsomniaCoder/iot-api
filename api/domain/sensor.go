package domain

import (
	"time"
)

type Sensor struct {
	ID           int64     `json:"id"`
	SoilMoisture float64   `json:"soil_moisture" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
}

type SensorUsecase interface {
	FetchAll() ([]Sensor, error)
	Store(*Sensor) error
}

type SensorRepository interface {
	FetchAll() ([]Sensor, error)
	Store(*Sensor) error
}
