package usecases

import (
	"context"
	"svc-order/app/models"
	"time"
)

type orderUsecase usecase

type OrderUsecase interface {
	CreateOrder(ctx context.Context, request models.CreateOrderRequest) (models.CreateOrderRequest, error)
}

func (u *orderUsecase) CreateOrder(ctx context.Context, request models.CreateOrderRequest) (models.CreateOrderRequest, error) {

	var (
		totalAmount float64
		response    models.CreateOrderRequest
		orderItem   models.OrderItem
	)

	for i, v := range request.OrderItem {
		product, err := u.Options.Repository.Product.GetDetailProduct(ctx, v.ProductId)
		if err != nil {
			return models.CreateOrderRequest{}, err
		}

		request.OrderItem[i].UnitPrice = product.Price
		totalAmount = totalAmount + product.Price
	}
	now := time.Now()

	orderReq := models.Order{
		UserId:          request.UserId,
		TotalAmount:     totalAmount,
		Status:          "pending",
		OrderDate:       now,
		PaymentDeadline: now.Add(5 * time.Minute),
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	tx := u.Options.Postgres.Begin()
	order, err := u.Options.Repository.Order.CreatOrder(tx, orderReq)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRequest{}, err
	}

	for _, v := range request.OrderItem {
		v.OrderId = order.OrderId
		v.CreatedAt = now
		v.UpdatedAt = now

		orderItem, err = u.Options.Repository.OrderItem.CreatOrderItem(tx, v)
		if err != nil {
			tx.Rollback()
			return models.CreateOrderRequest{}, err
		}
		response.OrderItem = append(response.OrderItem, orderItem)

		resvStockReq := models.ReservedStock{
			OrderItemId:           orderItem.OrderItemId,
			ProductId:             v.ProductId,
			ReservedQuantity:      v.Quantity,
			ReservationExpiryTime: order.PaymentDeadline,
			CreatedAt:             now,
			UpdatedAt:             now,
		}
		_, err = u.Options.Repository.ReservedStock.CreatReservedStock(tx, resvStockReq)
		if err != nil {
			tx.Rollback()
			return models.CreateOrderRequest{}, err
		}
	}

	tx.Commit()
	response.Order = order
	return response, nil
}
