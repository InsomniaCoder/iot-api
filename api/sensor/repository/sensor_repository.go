package repository

import (
	log "github.com/sirupsen/logrus"

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

func (s *sensorRepository) FetchAll() (sensorSlice *[]domain.Sensor, err error) {
	log.Println("sensor repository fetching all")

	var sensorFound []domain.Sensor = make([]domain.Sensor, 0)

	if err := s.con.Find(&sensorFound).Error; err != nil {
		return nil, err
	}

	log.Printf("%d rows found.\n", len(sensorFound))
	log.Printf("%v\n", sensorFound)
	return &sensorFound, nil
}

func (s *sensorRepository) Store(sensorData *domain.Sensor) (createdSensor *domain.Sensor, err error) {
	log.Printf("sensor repository saving %+v/n", sensorData)

	result := s.con.Create(sensorData)

	if result.Error != nil || result.RowsAffected != 1 {
		return nil, result.Error
	} else {
		log.Printf("sensor data created: %+v\n", sensorData)
		return sensorData, nil
	}
}
