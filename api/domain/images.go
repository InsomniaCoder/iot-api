package domain

import (
	"time"
)

type ImageCommand struct {
	CreatedAt time.Time `json:"created_at"`
}

type ImageLink struct {
	S3Url     string    `json:"s3Url"`
	CreatedAt time.Time `json:"created_at"`
}

type ImageUsecase interface {
	SearchImageByTimeRange(startTime time.Time, endTime time.Time) ([]ImageLink, error)
	SendCaptureCommand() error
}

type ImageRepository interface {
	FetchImage(startTime time.Time, endTime time.Time) ([]ImageLink, error)
	SendCaptureCommand() error
}
