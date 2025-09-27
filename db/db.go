package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS precious (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT,
		key TEXT
	)
	`)
	if err != nil {
		panic(err)
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
}
