migrate-up:
	migrate -path internal/db/migration/ -database "postgresql://postgres:postgres@localhost:5432/people?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migration/ -database "postgresql://postgres:postgres@localhost:5432/people?sslmode=disable" -verbose down