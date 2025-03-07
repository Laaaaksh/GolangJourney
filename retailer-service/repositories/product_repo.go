package repositories

import (
	"gorm.io/gorm"
	"retailer-service/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	GetAll() ([]models.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *productRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
