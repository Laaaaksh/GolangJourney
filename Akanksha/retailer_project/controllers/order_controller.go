package controllers

import (
	"fmt"
	"net/http"
	"retailer_project/config"
	"retailer_project/models"
	"time"

	"github.com/gin-gonic/gin"
)

const CooldownDuration = 5 * time.Minute // Cooldown period

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique cooldown key for the customer
	cooldownKey := fmt.Sprintf("cooldown:%s", order.CustomerID)

	// Step 1: Check if the cooldown key exists in Redis
	exists, err := config.RedisClient.Exists(config.Ctx, cooldownKey).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
		return
	}

	if exists > 0 {
		// If the cooldown key exists, the customer is still in cooldown
		c.JSON(http.StatusForbidden, gin.H{"error": "Cooldown period active. Try again later."})
		return
	}

	// Step 2: Place the order (since cooldown is not active)
	order.CreatedAt = time.Now()
	order.Status = "order placed"
	config.DB.Create(&order)

	// Step 3: Store cooldown in Redis with TTL (expires after 5 minutes)
	config.RedisClient.Set(config.Ctx, cooldownKey, "active", CooldownDuration)

	// Step 4: Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully!", "order": order})
}
