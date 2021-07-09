package usecase

import (
	"time"

	"github.com/insomniacoder/iot-api/api/domain"
	log "github.com/sirupsen/logrus"
)

type imageUsecase struct {
	imageRepository domain.ImageRepository
}

func NewImageUsecase(imageRepo domain.ImageRepository) domain.ImageUsecase {
	return &imageUsecase{
		imageRepository: imageRepo,
	}
}

func (i *imageUsecase) SearchImageByTimeRange(startTime time.Time, endTime time.Time) (*[]domain.ImageLink, error) {
	log.Printf("imageUsecase SearchImageByTimeRange %v , %v\n", startTime, endTime)
	return nil, nil
}

func (i *imageUsecase) SendCaptureCommand() error {
	log.Println("imageUsecase SendCaptureCommand")
	return i.imageRepository.SendCaptureCommand()
}
