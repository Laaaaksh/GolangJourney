package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"retailer-service/handlers"
	"retailer-service/middleware"
	"retailer-service/models"
	"retailer-service/repositories"
	"retailer-service/services"
)

func main() {
	// Connect to MySQL
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
	orderService := services.NewOrderService(orderRepo)

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
