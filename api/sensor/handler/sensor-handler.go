package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/iot-api/api/domain"

	log "github.com/sirupsen/logrus"
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
	var sensorData domain.Sensor

	c.Bind(&sensorData)

	log.Printf("%+v\n", sensorData)
}
