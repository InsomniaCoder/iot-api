package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/iot-api/api/domain"
)

type SensorHandler struct {
	SensorUsecase domain.SensorUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewSensorHandler(c *gin.Engine, sensorUsecase domain.SensorUsecase) {
	handler := &SensorHandler{
		SensorUsecase: sensorUsecase,
	}

	api := c.Group("api/v1")
	{
		api.POST("sensors", handler.CreateSensor)
	}
}

func (s *SensorHandler) CreateSensor(c *gin.Context) {
}
