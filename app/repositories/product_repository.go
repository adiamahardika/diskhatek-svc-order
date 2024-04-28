package repositories

import (
	"context"
	"svc-order/app/models"
)

type productRepository repository

type ProductRepository interface {
	GetDetailProduct(ctx context.Context, id int) (models.Product, error)
}

func (r *productRepository) GetDetailProduct(ctx context.Context, id int) (models.Product, error) {

	var (
		product models.Product
	)

	stockQuery := r.Options.Postgres.Table("stocks").Select("COALESCE(SUM(quantity),0)").Joins("JOIN warehouses ON stocks.warehouse_id = warehouses.warehouse_id AND warehouses.status = 'active'").Where("stocks.product_id = ?", id)
	reservedStockQuery := r.Options.Postgres.Table("reserved_stocks").Select("COALESCE(SUM(reserved_quantity),0)").Where("reserved_stocks.product_id = ?", id)

	query := r.Options.Postgres.Table("products").Select("products.*, shops.name AS shop, (?) - (?) AS available_stock", stockQuery, reservedStockQuery).Joins("JOIN shops ON products.shop_id = shops.shop_id").Where("products.product_id = ?", id)

	error := query.WithContext(ctx).Find(&product).Error

	return product, error
}
