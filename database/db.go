package database

import (
	"fmt"
	"log"

	"github.com/aryansh13/go_restapi_gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "3306"
	dbname   = "go_restapi_gin"
	db       *gorm.DB
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	fmt.Println("Connection success to database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
