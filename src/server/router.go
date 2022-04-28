package server

import (
	"github.com/Kelvedler/GoBlog/endpoints"
	"github.com/gin-gonic/gin"
)

func RootRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	v1 := api.Group("/v1")
	AuthRouter(v1.Group(""))
	UserRouter(v1.Group("/user"))
}

func AuthRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", endpoints.Register)
}

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", endpoints.List)
	routerGroup.GET("/:user_id", endpoints.Single)
	routerGroup.PUT("/:user_id", endpoints.Update)
	routerGroup.DELETE("/:user_id", endpoints.Delete)
}
