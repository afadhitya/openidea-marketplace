package productrepositories

import (
	"context"

	productentities "github.com/widcha/openidea-marketplace/internal/app/modules/product"
)

type IRepo interface {
	Product() IProduct
}

type IProduct interface {
	Create(context.Context, productentities.Product) error
}
