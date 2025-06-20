package utils

import (
	"fmt"
	"log"
	"snapdragon-details-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:0000@tcp(127.0.0.1:3306)/snapdb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database Connected!")
	DB.AutoMigrate(&models.User{})

}
