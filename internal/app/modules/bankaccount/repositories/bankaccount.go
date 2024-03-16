package bankaccountrepositories

import (
	"context"
	"log"

	bankaccountentities "github.com/widcha/openidea-marketplace/internal/app/modules/bankaccount"
)

type bankAccountRepo struct {
	repo *repo
}

func newBankAccountRepo(repo *repo) IBankAccount {
	return &bankAccountRepo{
		repo: repo,
	}
}

func (g *bankAccountRepo) Create(ctx context.Context, bankAccount bankaccountentities.BankAccount) error {
	query := `INSERT INTO bank_accounts (id, bank_name, bank_account_name, bank_account_number, user_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := g.repo.datasource.Postgre.ExecContext(ctx, query, bankAccount.Id, bankAccount.BankName, bankAccount.BankAccountName, bankAccount.BankAccountNumber, bankAccount.UserId)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
