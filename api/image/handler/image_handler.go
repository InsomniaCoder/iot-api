package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/iot-api/api/domain"

	log "github.com/sirupsen/logrus"
)

type ImageHandler struct {
	imageUsecase domain.ImageUsecase
}

func NewImageHandler(c *gin.Engine, iU domain.ImageUsecase) {
	handler := &ImageHandler{
		imageUsecase: iU,
	}

	api := c.Group("api/v1")
	{
		api.POST("images/search", handler.SearchImageByTime)
		api.POST("images/capture", handler.SendImageCaptureCommand)
	}
}

type ImageSearchRequest struct {
	startTime time.Time
	endTime   time.Time
}

func (i *ImageHandler) SearchImageByTime(c *gin.Context) {

	log.Println("ImageHandler SearchImageByTime")

	var searchRequest ImageSearchRequest

	c.Bind(&searchRequest)

	log.Printf("getting search request %+v\n", searchRequest)

	imageLinkSlice, err := i.imageUsecase.SearchImageByTimeRange(searchRequest.startTime, searchRequest.endTime)

	if err != nil {
		log.Panicf("search image fail: %+v\n", err)
	} else {
		c.JSON(200, imageLinkSlice)
	}

}

func (i *ImageHandler) SendImageCaptureCommand(c *gin.Context) {

	log.Println("ImageHandler SendImageCaptureCommand")

	if err := i.imageUsecase.SendCaptureCommand(); err != nil {
		log.Panicf("send image capture command failed: %+v\n", err)
	} else {
		c.JSON(200, "command sent")
	}
}
