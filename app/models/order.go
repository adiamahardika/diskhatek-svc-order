package models

import "time"

type Order struct {
	OrderId         int       `json:"order_id" gorm:"primaryKey"`
	UserId          int       `json:"user_id"`
	OrderDate       time.Time `json:"order_date"`
	TotalAmount     float64   `json:"total_amount"`
	Status          string    `json:"status"`
	PaymentDeadline time.Time `json:"payment_deadline"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
	Order
	OrderItem []OrderItem `json:"order_items"`
}
