package data_access

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

/*
	Holds postgreSQL db instance
*/
var _db *sql.DB;

/*
	Opens connection to DB and initialize it's instance
*/
func init() {
	// Take from heroku app engine, if missing take local DB
	connectionString := os.Getenv("DATABASE_URL");
	if (connectionString == "") {
		connectionString = "postgres://postgres:password@localhost/CodeChallenge?sslmode=disable"
	}

	db, err := sql.Open("postgres", connectionString)
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

