package bankaccountrepositories

import (
	"context"

	userbankaccountentities "github.com/widcha/openidea-marketplace/internal/app/modules/bankaccount"
)

type IRepo interface {
	BankAccount() IBankAccount
}

type IBankAccount interface {
	Create(context.Context, userbankaccountentities.BankAccount) error
}
