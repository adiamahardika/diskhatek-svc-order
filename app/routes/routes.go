package routes

import (
	"svc-order/app/controllers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, controller *controllers.Main) {
	v1 := e.Group("/v1")
	{
		order := v1.Group("/order")
		{
			order.POST("", controller.Order.Create)
			order.POST("/payment", controller.Order.PaymentOrder)
		}
	}
}
