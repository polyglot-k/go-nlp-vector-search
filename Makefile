# Makefile

.PHONY: start stop


start:
	docker-compose up --build -d
	go run cmd/main.go

stop:
	docker-compose down --remove-orphans
