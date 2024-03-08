GOPATH := $(shell go env GOPATH)
BINARY_NAME=parkwise
MIGRATE_BIN := $(GOPATH)/bin/migrate

DB_NAME := $(or $(DB_NAME),parkwise)
DB_USER := $(or $(DB_USER),postgres)
DB_PASS := $(or $(DB_PASSWORD),docker)
DB_HOST := $(or $(DB_HOST),localhost)
DB_PORT := $(or $(DB_PORT),5432)

$(MIGRATE_BIN):
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

$(MOCK_BIN):
	go install go.uber.org/mock/mockgen@latest

run:
	go run ./cmd/server/main.go

test:
	go test -v -cover -short ./...

swag:
	swag init -g ./cmd/server/main.go -o ./docs

clean:
	go clean

mock-db: $(MOCK_BIN)
	mockgen -package mockdb -destination internal/infrastructure/datastores/mockdb/parkingLot.go parkwise/internal/domain/definition ParkingLotRepository
	mockgen -package mockdb -destination internal/infrastructure/datastores/mockdb/vehicle.go parkwise/internal/domain/definition VehicleRepository

migrate-up: $(MIGRATE_BIN)
	migrate -source file://db/migrations -database postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable -verbose up

migrate-down: $(MIGRATE_BIN)
	migrate -source file://db/migrations -database postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable -verbose down

.PHONY: run test swag clean mock migrate-up migrate-down mock-db
