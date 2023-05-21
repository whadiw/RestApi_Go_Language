package database

import (
	"go-crud-restapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_wahyu"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Karyawan{})

	DB = database
}
