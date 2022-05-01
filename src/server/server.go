package server

import (
	"context"

	"github.com/Kelvedler/GoBlog/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Init(ctx context.Context, conn *pgx.Conn) {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(middlewares.TopContextMiddleware(ctx))
	engine.Use(middlewares.DBConnMiddleware(conn))
	RootRouter(engine)
	engine.Run()
}
