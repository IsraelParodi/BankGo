package main

import (
	"database/sql"
	"log"

	"github.com/israelparodi/bankgo/api/config"
	"github.com/israelparodi/bankgo/api/router"
	db "github.com/israelparodi/bankgo/db/sqlc/queries"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secretpostgres@localhost:5432/bank_go?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

var err error

func main() {
	config.DB, err = sql.Open(dbDriver, dbSource)
	config.Queries = db.New(config.DB)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer config.DB.Close()

	r := router.SetupRouter()
	err = r.Run()
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
