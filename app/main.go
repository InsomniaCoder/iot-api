package main

import (
	"fmt"

	"github.com/insomniacoder/iot-api/config"
)

func init() {
	fmt.Printf("initializaing main")
	fmt.Printf("loaded config: %+v", config.Config)
}

func main() {

}
