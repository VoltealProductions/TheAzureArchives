FROM golang:1.23.0-bookworm AS build

WORKDIR /app

COPY . .

RUN go mod download

# Build the API Application
RUN go build -o ./archives ./cmd

# Build the Migrator
RUN go build -o ./migrator ./cmd/migrate

# Copy the migrations
COPY ./cmd/migrate/migrations ./migrations

FROM debian:12.6 AS product

WORKDIR /app

COPY --from=build /app/archives .
COPY --from=build /app/migrator .
COPY --from=build /app/migrations ./migrations
COPY --from=build /app/.env .

EXPOSE 3030

CMD [ "/app/archives" ]