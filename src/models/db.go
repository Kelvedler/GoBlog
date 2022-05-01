package models

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func Init(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return conn, errors.New(fmt.Sprintf("Unable to connect to database %v\n", err))
	} else {
		return conn, nil
	}
}

func CloseDBConn(ctx context.Context, conn *pgx.Conn) {
	conn.Close(ctx)
}
