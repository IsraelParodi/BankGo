package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secretpostgres@localhost:5432/bank_go?sslmode=disable"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
