package repositories

import (
	"svc-order/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	Order         OrderRepository
	OrderItem     OrderItemRepository
	Product       ProductRepository
	ReservedStock ReservedStockRepository
	User          UserRepository
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
		Order:         (*orderRepository)(repo),
		OrderItem:     (*orderItemRepository)(repo),
		Product:       (*productRepository)(repo),
		ReservedStock: (*reservedStockRepository)(repo),
		User:          (*userRepository)(repo),
	}

	return m
}
