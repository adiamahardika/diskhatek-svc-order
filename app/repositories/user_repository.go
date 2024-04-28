package repositories

import (
	"context"
	"svc-order/app/models"
)

type userRepository repository

type UserRepository interface {
	GetUserDetail(ctx context.Context, id int) (models.User, error)
}

func (r *userRepository) GetUserDetail(ctx context.Context, id int) (models.User, error) {

	var (
		user models.User
	)

	query := r.Options.Postgres.Table("users").Where("users.user_id = ?", id)

	error := query.WithContext(ctx).Find(&user).Error

	return user, error
}
