package repositories

import (
	"svc-order/app/models"

	"gorm.io/gorm"
)

type reservedStockRepository repository

type ReservedStockRepository interface {
	CreatReservedStock(tx *gorm.DB, request models.ReservedStock) (models.ReservedStock, error)
}

func (r *reservedStockRepository) CreatReservedStock(tx *gorm.DB, request models.ReservedStock) (models.ReservedStock, error) {

	err := tx.Save(&request).Error

	return request, err
}
