package usecase

import (
	"github.com/insomniacoder/iot-api/api/domain"
	log "github.com/sirupsen/logrus"
)

type sensorUsecase struct {
	sensorRepository domain.SensorRepository
}

func NewSensorUsecase(sensorRepo domain.SensorRepository) domain.SensorUsecase {
	return &sensorUsecase{
		sensorRepository: sensorRepo,
	}
}

func (s *sensorUsecase) FetchAll() (sensorSlice *[]domain.Sensor, err error) {
	log.Println("fetching all sensor usecase...")

	sensorSlice, err = s.sensorRepository.FetchAll()

	if err != nil {
		return nil, err
	}
	return sensorSlice, nil
}

func (s *sensorUsecase) Store(sensorData *domain.Sensor) (savedSensor *domain.Sensor, err error) {

	log.Println("sensor storing usecase...")

	savedSensor, err = s.sensorRepository.Store(sensorData)

	if err != nil {
		return nil, err
	}
	return savedSensor, nil
}
