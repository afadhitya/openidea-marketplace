package userrepositories

import "github.com/widcha/openidea-marketplace/internal/pkg"

type repo struct {
	datasource *pkg.Datasource
}

func NewRepo(datasource *pkg.Datasource) IRepo {
	return &repo{
		datasource: datasource,
	}
}

func (r *repo) User() IUser {
	return newUserRepo(r)
}
