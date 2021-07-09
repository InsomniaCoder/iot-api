package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/insomniacoder/iot-api/api/domain"
	"github.com/insomniacoder/iot-api/config"
	"github.com/insomniacoder/iot-api/kafka/producer"
)

type imageRepository struct {
	commandTopic string
}

func NewImageRepository() domain.ImageRepository {
	topic := config.Config.Kafka.Producer.Topic
	return &imageRepository{topic}
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
	ctx := context.Background()
	//block send
	return producer.Produce(ctx, i.commandTopic, "capture")
}
