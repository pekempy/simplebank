package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pekempy/simplebank/api"
	db "github.com/pekempy/simplebank/db/sqlc"
	"github.com/pekempy/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server ", err)
	}
}