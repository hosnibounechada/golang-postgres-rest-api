migration_up: DB_USERNAME=postgres DB_PASSWORD=password DB_NAME=gindb migrate -path $(CURDIR)/migrations -database "postgres://$${DB_USERNAME}:$${DB_PASSWORD}@db:5432/$${DB_NAME}?sslmode=disable" -verbose up

# migration_up: migrate -path .\migrations -database "postgresql://postgres:password@localhost:5432/gindb?sslmode=disable" -verbose up

migration_down: migrate -path .\migrations -database "postgresql://postgres:password@localhost:5432/gindb?sslmode=disable" -verbose down

migration_fix: migrate -path .\migrations -database "postgresql://postgres:password@localhost:5432/gindb?sslmode=disable" force VERSION