package repositories

import (
	"svc-order/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	Order     OrderRepository
	OrderItem OrderItemRepository
	Product   ProductRepository
}

type repository struct {
	Options Options
}

type Options struct {
	Postgres *gorm.DB
	Config   *config.Config
}

func Init(opts Options) *Main {
	repo := &repository{opts}

	m := &Main{
		Order:     (*orderRepository)(repo),
		OrderItem: (*orderItemRepository)(repo),
		Product:   (*productRepository)(repo),
	}

	return m
}
