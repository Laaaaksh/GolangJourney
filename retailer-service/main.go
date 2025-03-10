package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"retailer-service/handlers"
	"retailer-service/middleware"

	"retailer-service/models"
	"retailer-service/repositories"
	"retailer-service/services"

	"github.com/go-redis/redis/v8"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Connect to MySQL
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// Ping Redis to verify connection
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	dsn := "root:Gymboi@10082001@tcp(127.0.0.1:3306)/go_gin_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate models
	db.AutoMigrate(&models.Product{}, &models.Order{}, &models.CustomerCoolDown{})

	// Initialize repositories
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Initialize services
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, redisClient)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Setup Gin router
	router := gin.Default()

	// Product routes
	router.POST("/product", productHandler.AddProduct)
	router.PATCH("/product/:id", productHandler.UpdateProduct)
	router.GET("/product/:id", productHandler.GetProduct)
	router.GET("/products", productHandler.GetAllProducts)

	// Order routes (protected by JWT middleware)
	router.POST("/order", middleware.AuthMiddleware(), orderHandler.PlaceOrder)
	router.GET("/order/:id", middleware.AuthMiddleware(), orderHandler.GetOrder)

	// Start server
	router.Run(":8080")
}
