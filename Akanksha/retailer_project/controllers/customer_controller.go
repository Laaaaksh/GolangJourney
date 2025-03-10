package controllers

import (
	"net/http"
	"retailer_project/config"
	"retailer_project/models"

	"github.com/gin-gonic/gin"
)

// CreateCustomer handles the creation of a new customer
func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	// Bind JSON input to customer struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save customer to the database
	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Customer created successfully",
		"customer": customer,
	})
}

// GetCustomer retrieves a single customer's details by ID
func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// GetAllCustomers retrieves all customers
func GetAllCustomers(c *gin.Context) {
	var customers []models.Customer

	if err := config.DB.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}
