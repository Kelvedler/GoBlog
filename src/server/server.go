package server

import "github.com/gin-gonic/gin"

func Init() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	RootRouter(engine)
	engine.Run()
}
