package services

import (
	"context"
	"errors"
	"fmt"
	"retailer-service/models"
	"retailer-service/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type OrderService interface {
	PlaceOrder(ctx context.Context, order *models.Order) error
	GetOrder(id string) (*models.Order, error)
}

type orderService struct {
	OrderRepo repositories.OrderRepository
	redis     *redis.Client
}

func NewOrderService(orderRepo repositories.OrderRepository, redisClient *redis.Client) OrderService {
	return &orderService{
		OrderRepo: orderRepo,
		redis:     redisClient,
	}
}

// Lua script for atomic cooldown check and update
var cooldownScript = redis.NewScript(`
local key = KEYS[1]
local current_time = tonumber(ARGV[1])
local cooldown = tonumber(ARGV[2])

local last_order = redis.call("GET", key)
if last_order then
    if current_time - tonumber(last_order) < cooldown then
        return 0
    end
end

redis.call("SET", key, current_time, "EX", cooldown)
return 1
`)

func (s *orderService) PlaceOrder(ctx context.Context, order *models.Order) error {
	// Execute Lua script atomically
	result, err := cooldownScript.Run(ctx, s.redis, []string{order.CustomerID},
		time.Now().Unix(), 300).Int()
	if err != nil {
		return fmt.Errorf("redis error: %v", err)
	}

	if result == 0 {
		return errors.New("customer is in cool-down period")
	}

	// Proceed with order creation
	return s.OrderRepo.Create(order)
}

func (s *orderService) GetOrder(id string) (*models.Order, error) {
	return s.OrderRepo.GetByID(id)
}
