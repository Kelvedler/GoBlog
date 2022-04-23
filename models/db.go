package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func Init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
	} else {
		db = conn
	}
}

func GetDB() *pgx.Conn {
	return db
}

func CloseDBConn() {
	db.Close(context.Background())
}
