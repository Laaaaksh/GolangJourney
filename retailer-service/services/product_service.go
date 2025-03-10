package services

import (
	"retailer-service/models"
	"retailer-service/repositories"
	"sync"
)

type ProductService interface {
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	GetProduct(id string) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
}

type productService struct {
	Repo  repositories.ProductRepository
	Mutex sync.Mutex
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{Repo: repo}
}

func (s *productService) AddProduct(product *models.Product) error {
	return s.Repo.Create(product)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Repo.Update(product)
}

func (s *productService) GetProduct(id string) (*models.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}
