package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectionDatabase() *sql.DB {
	connection := "user=root dbname=apigo password="" host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
