package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"retailer-service/handlers"
	"retailer-service/models"
	"testing"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) AddProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) UpdateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) GetProduct(id string) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) GetAllProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func TestAddProductHandler(t *testing.T) {
	mockService := new(MockProductService)
	mockService.On("AddProduct", mock.Anything).Return(nil)

	handler := handlers.NewProductHandler(mockService)

	// Create a Gin router
	router := gin.Default()
	router.POST("/product", handler.AddProduct)

	// Create a test request
	product := models.Product{
		ProductName: "bottle",
		Price:       50,
		Quantity:    40,
	}
	body, _ := json.Marshal(product)
	req, _ := http.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
