package main

import (
	"github.com/Kelvedler/GoBlog/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Init()
	GinEngine := gin.Default()
	GinEngine.Run()
}
