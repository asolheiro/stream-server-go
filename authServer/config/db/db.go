package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"host=postgres user=postgres password=password dbname= streamkeys sslmode=disable",
	)
	if err != nil {
		log.Fatal("failed to connect to database, err: ", err)
		return nil, err
	}

	err = db.Ping()
	return db, err
}
