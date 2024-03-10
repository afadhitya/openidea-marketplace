package pkg

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     string

	MaxIdleConn     int
	MaxOpenConn     int
	MaxLifeTimeConn int
}

func NewPostgres(config PostgresConfig) *sqlx.DB {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s", config.Username, config.DBName, config.Password, config.Host, config.Port)
	client, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
