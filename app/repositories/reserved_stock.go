package repositories

import (
	"svc-order/app/models"

	"gorm.io/gorm"
)

type reservedStockRepository repository

type ReservedStockRepository interface {
	CreateReservedStock(tx *gorm.DB, request models.ReservedStock) (models.ReservedStock, error)
	DeleteReservedStock(tx *gorm.DB, request models.ReservedStock) error
}

func (r *reservedStockRepository) CreateReservedStock(tx *gorm.DB, request models.ReservedStock) (models.ReservedStock, error) {

	err := tx.Save(&request).Error

	return request, err
}

func (r *reservedStockRepository) DeleteReservedStock(tx *gorm.DB, request models.ReservedStock) error {

	if request.OrderItemId != 0 {
		tx = tx.Where("order_item_id = ?", request.OrderItemId)
	}

	err := tx.Delete(&request).Error

	return err
}
