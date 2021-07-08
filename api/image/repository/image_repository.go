package repository

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/insomniacoder/iot-api/api/domain"
)

type imageRepository struct {
}

func NewSensorRepository() domain.ImageRepository {
	return &imageRepository{}
}

type ImageRepository interface {
	FetchImage(startTime time.Time, endTime time.Time) (*[]domain.ImageLink, error)
	SendCaptureCommand() error
}

func (i *imageRepository) FetchImage(startTime time.Time, endTime time.Time) (*[]domain.ImageLink, error) {
	log.Println("image repository fetch image")
	return nil, nil
}

func (i *imageRepository) SendCaptureCommand() error {
	log.Println("image repository send capture command")
	return nil
}
