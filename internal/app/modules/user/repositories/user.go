package userrepositories

import (
	"context"
	"log"

	userentities "github.com/widcha/openidea-marketplace/internal/app/modules/user"
)

type userRepo struct {
	repo *repo
}

func newUserRepo(repo *repo) IUser {
	return &userRepo{
		repo: repo,
	}
}

func (g *userRepo) Create(ctx context.Context, user userentities.User) error {
	query := `INSERT INTO users (id, name, username, password) VALUES ($1, $2, $3, $4)`
	_, err := g.repo.datasource.Postgre.ExecContext(ctx, query, user.Id, user.Name, user.Username, user.Password)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
