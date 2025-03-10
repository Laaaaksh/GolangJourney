package models


type Customer struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
