package repositories

import (
	"svc-order/app/models"

	"gorm.io/gorm"
)

type stockRepository repository

type StockRepository interface {
	ReduceStock(tx *gorm.DB, request models.Stock) error
}

func (r *stockRepository) ReduceStock(tx *gorm.DB, request models.Stock) error {

	stock := []models.Stock{}
	err := tx.Table("stocks").Select("stocks.*").Joins("JOIN warehouses ON warehouses.warehouse_id = stocks.warehouse_id AND warehouses.status = 'active'").Where("product_id = ?", request.ProductId).Order("stocks.quantity DESC").Find(&stock).Error
	if err != nil {
		return err
	}

	for _, v := range stock {

		if request.Quantity <= 0 {
			break
		}

		columns := map[string]interface{}{
			"updated_at": request.UpdatedAt,
		}

		if v.Quantity >= request.Quantity {
			columns["quantity"] = v.Quantity - request.Quantity
		} else if v.Quantity-request.Quantity < 0 {
			columns["quantity"] = 0
		}

		err = tx.Table("stocks").Select("quantity").Where("stocks.stock_id = ?", v.StockId).Updates(columns).Error
		if err != nil {
			return err
		}

		request.Quantity = request.Quantity - v.Quantity
	}

	return nil
}
