package controllers

import (
	"net/http"
	"svc-order/app/constants"
	"svc-order/app/helpers"
	"svc-order/app/models"

	customError "svc-order/pkg/customerrors"

	"github.com/ezartsh/inrequest"
	"github.com/ezartsh/validet"
	"github.com/labstack/echo/v4"
)

type orderController controller

type OrderController interface {
	Create(ctx echo.Context) error
}

func (c *orderController) Create(ctx echo.Context) error {
	var (
		reqBody models.CreateOrderRequest
		err     error
	)

	req, err := inrequest.Json(ctx.Request())
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	mapReq := req.ToMap()
	schema := validet.NewSchema(
		mapReq,
		map[string]validet.Rule{},
		validet.Options{},
	)

	errorBags, err := schema.Validate()
	if err != nil {
		err := customError.NewBadRequestError(err.Error())
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), errorBags.Errors, nil, nil)
	}

	err = req.ToBind(&reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	// err = c.Options.UseCases.Validate.IsValidCreateOrders(ctx.Request().Context(), mapReq)
	// if err != nil {
	// 	return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	// }
	// order, err = c.Options.UseCases.Order.CreateOrder(order)
	// if err != nil {
	// 	return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
	// }

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, reqBody, nil)
}