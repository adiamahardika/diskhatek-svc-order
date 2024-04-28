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

	query := r.Options.Postgres.Table("products").Select("products.*, shops.name AS shop, COALESCE(SUM(stocks.quantity),0) - COALESCE(SUM(reserved_stocks.reserved_quantity),0) AS available_stock").Joins("JOIN shops ON products.shop_id = shops.shop_id").Joins("JOIN stocks ON products.product_id = stocks.product_id").Joins("JOIN warehouses ON stocks.warehouse_id = warehouses.warehouse_id AND warehouses.status = 'active'").Joins("LEFT JOIN reserved_stocks ON products.product_id = reserved_stocks.product_id").Where("products.product_id = ?", id).Group("products.product_id, shops.name")

	error := query.WithContext(ctx).Find(&product).Error

	return product, error
}
