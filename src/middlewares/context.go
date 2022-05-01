package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

func TopContextMiddleware(topContext context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("topContext", topContext)
		ctx.Next()
	}
}
