package repositories

import (
	"context"
	"svc-order/app/models"

	"gorm.io/gorm"
)

type orderItemRepository repository

type OrderItemRepository interface {
	CreateOrderItem(tx *gorm.DB, request models.OrderItem) (models.OrderItem, error)
	GetOrderItem(ctx context.Context, request models.OrderItem) ([]models.OrderItem, error)
}

func (r *orderItemRepository) CreateOrderItem(tx *gorm.DB, request models.OrderItem) (models.OrderItem, error) {

	err := tx.Save(&request).Error

	return request, err
}

func (r *orderItemRepository) GetOrderItem(ctx context.Context, request models.OrderItem) ([]models.OrderItem, error) {
	var (
		orderItems []models.OrderItem
	)

	query := r.Options.Postgres.Table("order_items")

	if request.OrderId != 0 {
		query = query.Where("order_id = ?", request.OrderId)
	}

	err := query.WithContext(ctx).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
