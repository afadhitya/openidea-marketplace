package health

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type Usecase struct {
	postgres *sqlx.DB
}

type Response struct {
	Status map[string]string
}

func NewUsecase(postgres *sqlx.DB) Usecase {
	return Usecase{
		postgres: postgres,
	}
}

func (u Usecase) HealthCheck(ctx context.Context) (Response, error) {
	if err := u.postgres.PingContext(ctx); err != nil {
		log.Fatalln(ctx, "error while pinging to database", err)
		return Response{
			Status: map[string]string{"database": "database connection error"},
		}, err
	}

	return Response{
		Status: map[string]string{
			"database": "healthy",
		},
	}, nil
}
