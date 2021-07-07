package usecase

import (
	"github.com/insomniacoder/iot-api/api/domain"
)

type sensorUsecase struct {
	sensorRepository domain.SensorRepository
}

func NewSensorUsecase(sensorRepo domain.SensorRepository) domain.SensorUsecase {
	return &sensorUsecase{
		sensorRepository: sensorRepo,
	}
}

func (s *sensorUsecase) FetchAll() (sensorSlice []domain.Sensor, err error) {
	return
}

func (s *sensorUsecase) Store(sensorData *domain.Sensor) (err error) {
	return
}
