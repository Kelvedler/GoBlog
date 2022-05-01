package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func DBConnMiddleware(conn *pgx.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("dbConn", conn)
		ctx.Next()
	}
}
