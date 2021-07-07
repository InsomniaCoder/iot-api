package repository

import (
	"github.com/insomniacoder/iot-api/api/domain"
)

type sensorRepository struct {
}

func NewSensorRepository() domain.SensorRepository {
	return &sensorRepository{}
}

func (s *sensorRepository) FetchAll() (sensorSlice []domain.Sensor, err error) {
	return
}

func (s *sensorRepository) Store(sensorData *domain.Sensor) (err error) {
	return
}
