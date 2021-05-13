package main

import (
	"database/sql"
	"log"
	"simplebank-api/api"
	db "simplebank-api/db/sqlc"
	"simplebank-api/utils"

	_ "github.com/lib/pq"
)

var (
	config utils.Config
)

func init() {
	var err error

	config, err = utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}

func main() {
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
