package config

import (
	"database/sql"

	db "github.com/israelparodi/bankgo/db/sqlc/queries"
)

var DB *sql.DB
var Queries *db.Queries
