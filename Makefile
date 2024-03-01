createdb:
	docker exec -it postgres12 createdb --username=root --owner=root begho_db

dropdb:
	docker exec -it postgres12 dropdb begho_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/begho_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/begho_db?sslmode=disable" -verbose down

sqlc:
	../sqlc generate


.PHONY:createdb dropdb migrateup  migratedown  sqlc test server 