package saveproduct

import (
	"context"
	"log"

	"github.com/google/uuid"
	productentities "github.com/widcha/openidea-marketplace/internal/app/modules/product"
	productrepositories "github.com/widcha/openidea-marketplace/internal/app/modules/product/repositories"
)

type interactor struct {
	productRepo productrepositories.IRepo
}

func NewUsecase(productRepo productrepositories.IRepo) Inport {
	return &interactor{
		productRepo: productRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) error {
	bankAccount := productentities.Product{
		Id: uuid.NewString(),
	}

	err := i.productRepo.Product().Create(ctx, bankAccount)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
