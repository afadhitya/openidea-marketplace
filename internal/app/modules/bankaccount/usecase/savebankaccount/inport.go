package savebankaccount

import "context"

type Inport interface {
	Execute(context.Context, InportRequest) error
}

type InportRequest struct {
	BankName          string `json:"bankName" validate:"required"`
	BankAccountName   string `json:"bankAccountName" validate:"required"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required"`
	UserID            string
}
