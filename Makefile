postgres:
	docker run --name postgres17 -d -p 5432:5432 -e POSTGRES_PASSWORD=mypw -e POSTGRES_USER=root postgres:17.5-alpine

createdb:
	docker exec -it 1d87016dc967 createdb cajita

dropdb:
	docker exec -it 1d87016dc967 dropdb cajita

migrateup:
	migrate -path internal/db/migrations/ -database "postgresql://root:mypw@localhost:5432/cajita?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations/ -database "postgresql://root:mypw@localhost:5432/cajita?sslmode=disable" -verbose down

server:
	go run cmd/app/main.go

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
