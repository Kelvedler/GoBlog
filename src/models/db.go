package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func Init() {
	conn, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err == nil {
		err = conn.Ping()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
	} else {
		db = conn
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDBConn() {
	db.Close()
}
