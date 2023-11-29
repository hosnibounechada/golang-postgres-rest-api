#!/bin/sh

# sleep 5

migrate -path /app/migrations -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up

./main
