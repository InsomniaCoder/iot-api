package main

import (
	"github.com/insomniacoder/iot-api/config"
	log "github.com/sirupsen/logrus"

	_ "github.com/insomniacoder/iot-api/api"
)

func init() {
	log.Println("initializaing main")
	log.Printf("loaded config: %+v \n", config.Config)
	if config.Config.Debug {
		log.Println("Running in DEBUG mode")
	}
}

func main() {

}
