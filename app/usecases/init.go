package usecases

import (
	"svc-order/app/repositories"
	"svc-order/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	Order OrderUsecase
}

type usecase struct {
	Options Options
}

type Options struct {
	Repository *repositories.Main
	Config     *config.Config
	Postgres   *gorm.DB
}

func Init(opts Options) *Main {
	uscs := &usecase{opts}

	m := &Main{
		Order: (*orderUsecase)(uscs),
	}

	return m
}
