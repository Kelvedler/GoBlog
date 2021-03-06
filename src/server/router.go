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
	routerGroup.GET("/google_auth", endpoints.GoogleAuthURL)
	routerGroup.POST("/register", endpoints.UserRegister)
	routerGroup.GET("/jwt", endpoints.JWTToken)
}

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", endpoints.UserList)
	routerGroup.GET("/:user_id", endpoints.UserSingle)
	routerGroup.PUT("/:user_id", endpoints.UserUpdate)
	routerGroup.DELETE("/:user_id", endpoints.UserDelete)
}
