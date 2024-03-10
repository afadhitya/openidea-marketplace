package pkg

import (
	"github.com/jmoiron/sqlx"
	"github.com/widcha/openidea-marketplace/configs"
)

type Datasource struct {
	Postgre *sqlx.DB
}

func NewDataSource() *Datasource {
	postgresClient := NewPostgres(PostgresConfig{
		Username: configs.Get().DBUsername,
		Password: configs.Get().DBPassword,
		DBName:   configs.Get().DBName,
		Host:     configs.Get().DBHost,
		Port:     configs.Get().DBPort,
	})

	return &Datasource{
		Postgre: postgresClient,
	}
}
