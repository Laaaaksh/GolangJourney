package services

import (
	"errors"
	"retailer-service/models"
	"retailer-service/repositories"
	"sync"
	"time"
)

type OrderService interface {
	PlaceOrder(order *models.Order) error
	GetOrder(id string) (*models.Order, error)
}

type orderService struct {
	OrderRepo repositories.OrderRepository
	CoolDown  map[string]int64
	Mutex     sync.Mutex
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderService{
		OrderRepo: orderRepo,
		CoolDown:  make(map[string]int64),
	}
}

func (s *orderService) PlaceOrder(order *models.Order) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	lastOrderTime, exists := s.CoolDown[order.CustomerID]
	if exists && time.Now().Unix()-lastOrderTime < 300 { // 5 minutes cool-down
		return errors.New("customer is in cool-down period")
	}

	s.CoolDown[order.CustomerID] = time.Now().Unix()
	return s.OrderRepo.Create(order)
}

func (s *orderService) GetOrder(id string) (*models.Order, error) {
	return s.OrderRepo.GetByID(id)
}
