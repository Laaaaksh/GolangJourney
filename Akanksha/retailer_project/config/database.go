package config

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"retailer_project/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Limcee@89@tcp(127.0.0.1:3306)/retailer_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	database.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{})
	DB = database
}
