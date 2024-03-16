package bankaccountrepositories

import "github.com/widcha/openidea-marketplace/internal/pkg"

type repo struct {
	datasource *pkg.Datasource
}

func NewRepo(datasource *pkg.Datasource) IRepo {
	return &repo{
		datasource: datasource,
	}
}

func (r *repo) BankAccount() IBankAccount {
	return newBankAccountRepo(r)
}
