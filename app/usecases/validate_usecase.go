package usecases

import (
	"context"
	"svc-order/app/models"
	customErrors "svc-order/pkg/customerrors"
)

type validateUsecase usecase

type ValidateUsecase interface {
	IsValidCreateOrder(ctx context.Context, request models.CreateOrderRequest) error
	IsValidPayment(ctx context.Context, request models.CreateOrderRequest) error
}

func (u *validateUsecase) IsValidCreateOrder(ctx context.Context, request models.CreateOrderRequest) error {

	user, err := u.Options.Repository.User.GetUserDetail(ctx, request.UserId)
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if user.UserId == 0 {
		return customErrors.NewBadRequestErrorf("User id %d not found", request.UserId)
	}

	for _, v := range request.OrderItem {
		product, err := u.Options.Repository.Product.GetDetailProduct(ctx, v.ProductId)
		if err != nil {
			return customErrors.NewInternalServiceError(err.Error())
		}
		if product.ProductId == 0 {
			return customErrors.NewBadRequestErrorf("Product id %d not found", v.ProductId)
		}
		if v.Quantity > product.AvailableStock {
			return customErrors.NewBadRequestErrorf("Insufficient stock available for product id %d", v.ProductId)
		}
	}

	return nil
}

func (u *validateUsecase) IsValidPayment(ctx context.Context, request models.CreateOrderRequest) error {

	order, err := u.Options.Repository.Order.GetOrderDetail(ctx, request.OrderId)
	if err != nil {
		return customErrors.NewInternalServiceError(err.Error())
	}
	if order.OrderId == 0 {
		return customErrors.NewBadRequestErrorf("Order id %d not found", request.OrderId)
	}
	if order.Status != "pending" {
		return customErrors.NewBadRequestErrorf("Payment order id %d cannot be made. Order status is %s", request.OrderId, order.Status)
	}

	return nil
}
