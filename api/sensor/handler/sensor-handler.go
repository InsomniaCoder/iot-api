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
		api.GET("sensors", handler.GetAllSensorData)
	}
}

func (s *SensorHandler) CreateSensor(c *gin.Context) {
	var sensorData domain.Sensor

	c.Bind(&sensorData)

	log.Printf("%+v\n", sensorData)

	savedSensor, err := s.SensorUsecase.Store(&sensorData)

	if err != nil {
		log.Panicf("store sensor data fail: %+v\n", err)
	} else {
		c.JSON(200, savedSensor)
	}

}

func (s *SensorHandler) GetAllSensorData(c *gin.Context) {

	log.Println("Get All Sensor Data")

	sensorSlices, err := s.SensorUsecase.FetchAll()

	if err != nil {
		log.Panicf("getting all sensor data fail: %+v\n", err)
	} else {
		c.JSON(200, sensorSlices)
	}

}
