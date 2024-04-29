package controllers

import (
	"svc-order/app/usecases"
	"svc-order/pkg/config"
)

type Main struct {
	Order OrderController
	User  UserController
}

type controller struct {
	Options Options
}

type Options struct {
	Config   *config.Config
	UseCases *usecases.Main
}

func Init(opts Options) *Main {
	ctrlr := &controller{opts}

	m := &Main{
		Order: (*orderController)(ctrlr),
		User:  (*userController)(ctrlr),
	}

	return m
}
