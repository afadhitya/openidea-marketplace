package userrepositories

import (
	"context"
	"database/sql"
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

func (g *userRepo) GetbyUsername(ctx context.Context, username string) (userentities.User, error) {
	var user userentities.User
	query := `SELECT * FROM users WHERE username = $1`

	err := g.repo.datasource.Postgre.GetContext(ctx, &user, query, username)
	if err != nil {
		if err != sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}
