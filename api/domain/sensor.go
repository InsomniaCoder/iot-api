package domain

import (
	"time"
)

type Sensor struct {
	ID           uint      `gorm:"primaryKey" json:"id,omitempty"`
	SoilMoisture float64   `json:"soil_moisture" binding:"required"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Sensor) TableName() string {
	return "sensors"
}

type SensorUsecase interface {
	FetchAll() (*[]Sensor, error)
	Store(*Sensor) (*Sensor, error)
}

type SensorRepository interface {
	FetchAll() (*[]Sensor, error)
	Store(*Sensor) (*Sensor, error)
}
