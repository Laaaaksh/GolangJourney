package models

import "time"

type Order struct {
	ID         string    `gorm:"primaryKey"`
	ProductID string `gorm:"type:varchar(255)"`
	CustomerID string `gorm:"type:varchar(255)"`
	Quantity   int
	Status     string
	CreatedAt  time.Time
}
