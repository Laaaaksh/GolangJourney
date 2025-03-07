package models

type CustomerCoolDown struct {
	CustomerID    string `gorm:"primaryKey"`
	LastOrderTime int64  `json:"last_order_time"`
}
