getpostgres:
	docker pull postgres:16-alpine

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpostgres -d postgres:16-alpine

runcontainer:
	docker start postgres16

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root bank_go

dropdb:
	docker exec -it postgres16 dropdb bank_go

migrateup:
	migrate -path db/migration -database "postgresql://root:secretpostgres@localhost:5432/bank_go?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secretpostgres@localhost:5432/bank_go?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go clean -testcache && go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb server
