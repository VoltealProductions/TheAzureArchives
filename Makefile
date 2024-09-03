build:
	go build -o bin/archives ./cmd/main.go

api-run: build
	./bin/archives

client-build:
	go build -o bin/client ./cmd/client/main.go

client-run: client-build
	./bin/client

test:
	go test -v ./... -count=1

tidy:
	go mod tidy

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down