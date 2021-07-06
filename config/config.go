package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Debug    string
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var (
	Config AppConfig
)

func init() {
	fmt.Printf("initializaing config")
	LoadConfiguration(&Config)
}

func LoadConfiguration(appConfig *AppConfig) {
	viper.SetConfigFile(`config.yaml`)
	err := viper.ReadInConfig()

	// Read Environment Variables
	viper.AutomaticEnv()
	if err != nil {
		fmt.Printf("Cannot read configuration , %v", err)
		panic(err)
	}

	marshalErr := viper.Unmarshal(appConfig)

	if marshalErr != nil {
		fmt.Printf("Unable to decode into struct, %v", marshalErr)
		panic(marshalErr)
	}
}
