package main

import (
	"context"

	"github.com/Kelvedler/GoBlog/models"
	"github.com/Kelvedler/GoBlog/server"
)

func main() {
	ctx := context.Background()
	conn, err := models.Init(ctx)
	if err != nil {
		panic(err)
	}
	defer models.CloseDBConn(ctx, conn)
	server.Init(ctx, conn)
}
