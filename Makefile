

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs


migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://jahongir:00@localhost:5432/h_database?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://jahongir:00@localhost:5432/h_database?sslmode=disable' down
