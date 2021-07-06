package main

import (
	"fmt"

	"github.com/insomniacoder/iot-api/config"
)

func main() {
	fmt.Printf("init")
	fmt.Printf("%+v", config.Config)
}
