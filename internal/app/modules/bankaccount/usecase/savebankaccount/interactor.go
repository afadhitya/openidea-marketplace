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

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	bankAccount := bankaccountentities.BankAccount{
		Id:                uuid.NewString(),
		BankName:          req.BankName,
		BankAccountName:   req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
	}

	err := i.bankAccountRepo.BankAccount().Create(ctx, bankAccount)
	if err != nil {
		log.Println(err)
		return InportResponse{}, err
	}

	return InportResponse{
		BankAccountId:     bankAccount.Id,
		BankName:          bankAccount.BankName,
		BankAccountName:   bankAccount.BankAccountName,
		BankAccountNumber: bankAccount.BankAccountNumber,
	}, nil
}
