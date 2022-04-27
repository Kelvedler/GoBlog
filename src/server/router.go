package server

import (
	"github.com/Kelvedler/GoBlog/endpoints"
	"github.com/gin-gonic/gin"
)

func RootRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	v1 := api.Group("/v1")
	UserRouter(v1.Group("/user"))
}

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", endpoints.List)
	routerGroup.POST("/register", endpoints.Register)
}
