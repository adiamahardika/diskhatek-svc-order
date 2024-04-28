package repositories

import (
	"svc-order/app/models"

	"gorm.io/gorm"
)

type orderRepository repository

type OrderRepository interface {
	CreatOrder(tx *gorm.DB, request models.Order) (models.Order, error)
}

func (r *orderRepository) CreatOrder(tx *gorm.DB, request models.Order) (models.Order, error) {

	err := tx.Save(&request).Error

	return request, err
}
