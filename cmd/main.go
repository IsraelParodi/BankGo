package main

import (
	"database/sql"
	"log"

	"github.com/israelparodi/bankgo/api/router"
	"github.com/israelparodi/bankgo/config"
	db "github.com/israelparodi/bankgo/db/sqlc/queries"
	_ "github.com/lib/pq"
)

func main() {
	environment, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}
	config.DB, err = sql.Open(environment.DBDriver, environment.DBSource)
	config.Queries = db.New(config.DB)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer config.DB.Close()

	r := router.SetupRouter()
	err = r.Run(environment.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
