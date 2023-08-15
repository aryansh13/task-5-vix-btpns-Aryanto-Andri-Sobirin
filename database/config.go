package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabase() (*gorm.DB, error) {
	//konfigurasi database

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
