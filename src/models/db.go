package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type dbConnection pgx.Conn

var db *pgx.Conn

func Init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
	} else {
		db = conn
	}
}

func CloseDBConn() {
	db.Close(context.Background())
}
