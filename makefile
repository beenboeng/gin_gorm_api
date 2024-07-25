.Phony: up down 

up: 
	go run ./migration/migrate.go

down:


build:
	go build -o ./bin/go_api ./cmd/main.go