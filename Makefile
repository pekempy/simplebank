postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:4fglgUdyq9vY3vM55u1M@simple-bank.co0qppxnk4tx.eu-west-2.rds.amazonaws.com/simple_bank" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:4fglgUdyq9vY3vM55u1M@simple-bank.co0qppxnk4tx.eu-west-2.rds.amazonaws.com/simple_bank" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:4fglgUdyq9vY3vM55u1M@simple-bank.co0qppxnk4tx.eu-west-2.rds.amazonaws.com/simple_bank" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:4fglgUdyq9vY3vM55u1M@simple-bank.co0qppxnk4tx.eu-west-2.rds.amazonaws.com/simple_bank" -verbose down 1

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/pekempy/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock migrateup1 migratedown1 