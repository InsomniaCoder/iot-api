package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func init() {

	fmt.Println("Initializing datbase connection")

	con := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.Database.User,
		Config.Database.Password, Config.Database.Host, Config.Database.Port,
		Config.Database.Name)

	fmt.Printf("Connection String is: %v", con)

	database, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		fmt.Printf("Unable to connect to database, %v", err)
		panic(err)
	}
	DBConnection = database
}
