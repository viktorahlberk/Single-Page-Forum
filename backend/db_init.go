package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb() *sql.DB {
	db, err := sql.Open("sqlite3", "SPADB.db")
	CheckErr(err)
	return db
}
