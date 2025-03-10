package controllers

import (
	"net/http"
	"retailer_project/config"
	"retailer_project/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")
    if err := config.DB.First(&product, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    config.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}