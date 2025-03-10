package routes

import (
	"github.com/gin-gonic/gin"
	"retailer_project/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Product routes
	r.POST("/products", controllers.CreateProduct)
	r.PATCH("/product/:id", controllers.UpdateProduct)
	r.GET("/product/:id", controllers.GetProduct)
	r.GET("/products", controllers.GetAllProducts)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	// Order routes
	r.POST("/orders", controllers.PlaceOrder)


	// Add Customer routes
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customer/:id", controllers.GetCustomer)
	r.GET("/customers", controllers.GetAllCustomers)

	return r
}
