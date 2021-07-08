package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/insomniacoder/iot-api/config"
	"github.com/insomniacoder/iot-api/database"

	//sensors
	_sensorHandler "github.com/insomniacoder/iot-api/api/sensor/handler"
	_sensorRepository "github.com/insomniacoder/iot-api/api/sensor/repository"
	_sensorUsecase "github.com/insomniacoder/iot-api/api/sensor/usecase"

	//images
	_imageHandler "github.com/insomniacoder/iot-api/api/image/handler"
	_imageRepository "github.com/insomniacoder/iot-api/api/image/repository"
	_imageUsecase "github.com/insomniacoder/iot-api/api/image/usecase"
	log "github.com/sirupsen/logrus"
)

func init() {

	log.Println("Initializting API package")

	//set up router
	r := gin.Default()
	r.Use(Cors())

	//set up sensor dependency
	sensorRepo := _sensorRepository.NewSensorRepository(database.DBConnection)
	sensorUsecase := _sensorUsecase.NewSensorUsecase(sensorRepo)
	_sensorHandler.NewSensorHandler(r, sensorUsecase)
	//set up image dependency
	imageRepo := _imageRepository.NewImageRepository()
	imageUsecase := _imageUsecase.NewImageUsecase(imageRepo)
	_imageHandler.NewImageHandler(r, imageUsecase)
	//start server
	portNumber := fmt.Sprintf(":%d", config.Config.Server.Port)
	r.Run(portNumber)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
