package repositories

import (
	"gorm.io/gorm"
	"retailer-service/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{DB: db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *orderRepository) GetByID(id string) (*models.Order, error) {
	var order models.Order
	err := r.DB.First(&order, "id = ?", id).Error
	return &order, err
}
