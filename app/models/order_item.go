package models

import "time"

type OrderItem struct {
	OrderItemId int       `json:"order_item_id"`
	OrderId     int       `json:"order_id"`
	ProductId   int       `json:"product_id"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unit_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
