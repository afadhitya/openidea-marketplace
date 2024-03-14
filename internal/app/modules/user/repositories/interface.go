package userrepositories

import (
	"context"

	userentities "github.com/widcha/openidea-marketplace/internal/app/modules/user"
)

type IRepo interface {
	User() IUser
}

type IUser interface {
	Create(context.Context, userentities.User) error
	GetbyUsername(ctx context.Context, username string) (userentities.User, error)
}
