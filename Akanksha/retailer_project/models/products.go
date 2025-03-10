package models

type Product struct {
	ID          string `gorm:"primaryKey"`
	ProductName string
	Price       int
	Quantity    int
}
