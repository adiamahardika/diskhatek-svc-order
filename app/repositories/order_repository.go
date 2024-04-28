package repositories

import (
	"context"
	"svc-order/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository repository

type OrderRepository interface {
	CreateOrder(tx *gorm.DB, request models.Order) (models.Order, error)
	UpdateOrder(tx *gorm.DB, request models.Order, id int) (models.Order, error)
	GetOrderDetail(ctx context.Context, id int) (models.Order, error)
}

func (r *orderRepository) CreateOrder(tx *gorm.DB, request models.Order) (models.Order, error) {

	err := tx.Save(&request).Error

	return request, err
}

func (r *orderRepository) UpdateOrder(tx *gorm.DB, request models.Order, id int) (models.Order, error) {

	var (
		err   error
		order models.Order
	)

	columns := map[string]interface{}{
		"updated_at": request.UpdatedAt,
	}

	if request.Status != "" {
		columns["status"] = request.Status
	}

	err = tx.Table("orders").Model(&order).Clauses(clause.Returning{}).Where("orders.order_id = ?", id).Updates(columns).Error

	return order, err
}

func (r *orderRepository) GetOrderDetail(ctx context.Context, id int) (models.Order, error) {

	var (
		order models.Order
	)

	query := r.Options.Postgres.Table("orders").Where("orders.order_id = ?", id)

	error := query.WithContext(ctx).Find(&order).Error

	return order, error
}
