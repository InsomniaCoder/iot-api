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

func (s *sensorUsecase) FetchAll() (sensorSlice []domain.Sensor, err error) {
	return
}

func (s *sensorUsecase) Store(sensorData *domain.Sensor) (err error) {

	log.Println("sensor storing usecase...")

	if err := s.sensorRepository.Store(sensorData); err != nil {
		return err
	}
	return nil
}
