package database

import (
	"fmt"

	"github.com/insomniacoder/iot-api/config"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func init() {

	log.Println("Initializing datbase connection")

	dbConfig := config.Config.Database

	con := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.User, dbConfig.Password,
		dbConfig.Host, dbConfig.Port, dbConfig.Name)

	log.Printf("Connection String is: %v \n", con)

	database, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		log.Printf("Unable to connect to database, %v \n", err)
		panic(err)
	}

	DBConnection = database
}
