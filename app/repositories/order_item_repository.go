package repositories

import (
	"svc-order/app/models"

	"gorm.io/gorm"
)

type orderItemRepository repository

type OrderItemRepository interface {
	CreatOrderItem(tx *gorm.DB, request models.OrderItem) (models.OrderItem, error)
}

func (r *orderItemRepository) CreatOrderItem(tx *gorm.DB, request models.OrderItem) (models.OrderItem, error) {

	err := tx.Save(&request).Error

	return request, err
}
