package savebankaccount

import (
	"context"
	"log"

	"github.com/google/uuid"
	bankaccountentities "github.com/widcha/openidea-marketplace/internal/app/modules/bankaccount"
	bankaccountrepositories "github.com/widcha/openidea-marketplace/internal/app/modules/bankaccount/repositories"
)

type interactor struct {
	bankAccountRepo bankaccountrepositories.IRepo
}

func NewUsecase(bankAccountRepo bankaccountrepositories.IRepo) Inport {
	return &interactor{
		bankAccountRepo: bankAccountRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) error {
	bankAccount := bankaccountentities.BankAccount{
		Id:                uuid.NewString(),
		BankName:          req.BankName,
		BankAccountName:   req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
		UserId:            req.UserID,
	}

	err := i.bankAccountRepo.BankAccount().Create(ctx, bankAccount)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
