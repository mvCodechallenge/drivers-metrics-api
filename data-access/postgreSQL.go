package data_access

import (
	"database/sql"
	_ "github.com/lib/pq"
)

/*
	Holds postgreSQL db instance
*/
var _db *sql.DB;

/*
	Opens connection to DB and initialize it's instance
*/
func init() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/CodeChallenge?sslmode=disable")
	if (err != nil) {
		panic(err)
	}

	// Check that connection to DB is alive
	err = db.Ping()
	if (err != nil) {
		panic(err)
	}

	_db = db
}

