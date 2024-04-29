package controllers

import (
	"svc-order/app/helpers"

	customError "svc-order/pkg/customerrors"

	"github.com/labstack/echo/v4"
)

type userController controller

type UserController interface {
	Authentication() echo.MiddlewareFunc
}

func (c *userController) Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ctx.Request().Header.Get("token")
			err := c.Options.UseCases.User.Authentication(token)
			if err != nil {
				return helpers.StandardResponse(ctx, customError.GetStatusCode(err), []string{err.Error()}, nil, nil)
			}
			return next(ctx)
		}
	}
}
