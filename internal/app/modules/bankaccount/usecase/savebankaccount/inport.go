package savebankaccount

import "context"

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	BankName          string `json:"bankName" validate:"required"`
	BankAccountName   string `json:"bankAccountName" validate:"required"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required"`
}

type InportResponse struct {
	BankAccountId     string
	BankName          string
	BankAccountName   string
	BankAccountNumber string
}
