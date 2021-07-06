package model

import "time"

type Sensor struct {
	ID           uint32    `json:"id"`
	SoilMoisture float32   `json:"soil_moisture"`
	CreatedAt    time.Time `json:"created_at"`
}
