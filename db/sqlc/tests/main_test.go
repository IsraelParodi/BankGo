package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/israelparodi/bankgo/config"
	db "github.com/israelparodi/bankgo/db/sqlc/queries"

	_ "github.com/lib/pq"
)

var testDB *sql.DB
var testQueries *db.Queries

func TestMain(m *testing.M) {
	environment, err := config.LoadConfig("../../..")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}
	testDB, err = sql.Open(environment.DBDriver, environment.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
