package models

import "time"

type Product struct {
	ProductId      int       `json:"product_id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          float64   `json:"price"`
	AvailableStock int       `json:"available_stock"`
	ShopId         int       `json:"shop_id"`
	Shop           string    `json:"shop,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
