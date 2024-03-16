include .env

dev:
	nodemon --exec go run main.go start --signal SIGTERM

migration-up:

	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -path db/migrations -verbose up

migration-down:
	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -path db/migrations -verbose down

migration-fix:
	migrate -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -path db/migrations force $(version)

migration-add:
	migrate create -ext sql -dir db/migrations -seq $(name)