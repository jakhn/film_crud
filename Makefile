

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs


migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://jakh:00@localhost:5432/film?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://jakh:00@localhost:5432/film?sslmode=disable' down
