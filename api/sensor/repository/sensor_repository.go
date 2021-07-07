package repository

import (
	"log"

	"github.com/insomniacoder/iot-api/api/domain"
	"gorm.io/gorm"
)

type sensorRepository struct {
	con *gorm.DB
}

func NewSensorRepository(dbConnection *gorm.DB) domain.SensorRepository {
	return &sensorRepository{
		con: dbConnection,
	}
}

func (s *sensorRepository) FetchAll() (sensorSlice []domain.Sensor, err error) {
	return
}

func (s *sensorRepository) Store(sensorData *domain.Sensor) (err error) {
	result := s.con.Create(sensorData)
	if result.Error != nil || result.RowsAffected != 1 {
		return result.Error
	} else {
		log.Printf("sensor data created: %+v\n", sensorData)
		return nil
	}
}
