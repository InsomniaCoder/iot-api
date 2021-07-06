package main

import (
	"fmt"

	"github.com/insomniacoder/iot-api/api/route"
	"github.com/insomniacoder/iot-api/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Println("initializaing main")
	log.Printf("loaded config: %+v \n", config.Config)
	if config.Config.Debug {
		log.Println("Running in DEBUG mode")
	}
}

func main() {
	//set up router
	r := route.SetupRouter()
	portNumber := fmt.Sprintf(":%d", config.Config.Server.Port)
	r.Run(portNumber)
}
