package usecases

import (
	"svc-order/app/models"
	"svc-order/pkg/customerrors"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type userUsecase usecase

type UserUsecase interface {
	Authentication(tokenString string) error
}

func (u *userUsecase) Authentication(tokenString string) error {

	claims := &models.Claims{}
	jwtKey := []byte(viper.GetString("TOKEN_SECRET"))
	token, error := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	validatorError, _ := error.(*jwt.ValidationError)
	if token == nil {
		return customerrors.NewBadRequestError("Please provide token!")
	} else if validatorError != nil && validatorError.Errors == jwt.ValidationErrorExpired {
		return customerrors.NewBadRequestError("Your token expired!")
	} else if error != nil {
		return customerrors.NewBadRequestError("Your token invalid!")
	}

	return nil
}
