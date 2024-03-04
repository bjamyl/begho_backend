package main

import (
	"database/sql"
	"log"

	"github.com/bjamyl/begho/api"
	db "github.com/bjamyl/begho/db/sqlc"
)

const (
	db_driver = "postgres"
	db_source = "postgresql://root:secret@localhost:5433/begho_db?sslmode=disable"
)

func main() {
	conn, err := sql.Open(db_driver, db_source)
	if err != nil {
		log.Fatal("could not make connection", err.Error())
		return
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	server.Run("0.0.0.0:5435")
}
