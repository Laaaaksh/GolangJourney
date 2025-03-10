package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"retailer-service/models"
	"retailer-service/services"
)

type OrderHandler struct {
	Service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if order.CustomerID == "" || order.ProductID == "" || order.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id, product_id, and quantity are required"})
		return
	}

	// Generate a unique ID for the order (e.g., "ORD1234")
	order.ID = "ORD_" + uuid.New().String()[:8] // Example: ORD_5e6f7g8h

	// Set default status if empty
	if order.Status == "" {
		order.Status = "order placed"
	}

	// Create a context from the Gin context
	ctx := c.Request.Context()

	// Call the service with the context
	if err := h.Service.PlaceOrder(ctx, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         order.ID,
		"product_id": order.ProductID,
		"quantity":   order.Quantity,
		"status":     order.Status,
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.Service.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
