//package tests
//
//import (
//	"bytes"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"net/http"
//	"net/http/httptest"
//	"retailer-service/handlers"
//	"retailer-service/models"
//	"testing"
//)
//
//type MockOrderService struct {
//	mock.Mock
//}
//
//func (m *MockOrderService) PlaceOrder(order *models.Order) error {
//	args := m.Called(order)
//	return args.Error(0)
//}
//
//func (m *MockOrderService) GetOrder(id string) (*models.Order, error) {
//	args := m.Called(id)
//	return args.Get(0).(*models.Order), args.Error(1)
//}
//
//func TestPlaceOrderHandler(t *testing.T) {
//	mockService := new(MockOrderService)
//	mockService.On("PlaceOrder", mock.Anything).Return(nil)
//
//	handler := handlers.NewOrderHandler(mockService)
//
//	// Create a Gin router
//	router := gin.Default()
//	router.POST("/order", handler.PlaceOrder)

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"retailer-service/handlers"
	"retailer-service/models"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) PlaceOrder(ctx context.Context, order *models.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockOrderService) GetOrder(id string) (*models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func TestPlaceOrderHandler(t *testing.T) {
	mockService := new(MockOrderService)
	mockService.On("PlaceOrder", mock.Anything, mock.Anything).Return(nil)

	handler := handlers.NewOrderHandler(mockService)

	// Create a Gin router
	router := gin.Default()
	router.POST("/order", handler.PlaceOrder)

	// Create a test request
	order := models.Order{
		CustomerID: "CST123",
		ProductID:  "PROD123",
		Quantity:   1,
	}
	body, _ := json.Marshal(order)
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

//
//	// Create a test request
//	order := models.Order{
//		CustomerID: "CST123",
//		ProductID:  "PROD123",
//		Quantity:   1,
//	}
//	body, _ := json.Marshal(order)
//	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//
//	// Record the response
//	w := httptest.NewRecorder()
//	router.ServeHTTP(w, req)
//
//	// Assert the response
//	assert.Equal(t, http.StatusOK, w.Code)
//	mockService.AssertExpectations(t)
//}
