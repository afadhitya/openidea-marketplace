package bankaccountentities

type BankAccount struct {
	Id                string `json:"id" db:"id"`
	BankName          string `json:"bank_name" db:"bank_name"`
	BankAccountName   string `json:"bank_account_name" db:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number" db:"bank_account_number"`
	UserId            string `json:"user_id" db:"user_id"`
}
