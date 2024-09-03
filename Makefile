build:
	go build -o bin/archives ./cmd/main.go

run: build
	./bin/archives

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

watch:
	@clear && tailwindcss -i ./public/css/input.css -o ./public/css/style.css --watch

minify:
	@tailwindcss -i ./public/css/input.css -o ./public/css/style.min.css --minify