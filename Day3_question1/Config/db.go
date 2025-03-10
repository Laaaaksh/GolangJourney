package config

import (
	"apicall/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Gymboi@10082001@tcp(127.0.0.1:3306)/go_gin_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected!")
	DB = db
}

func MigrateDatabase() {
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrated!")
}
