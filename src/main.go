package main

import (
	"github.com/Kelvedler/GoBlog/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Init()
	defer models.CloseDBConn()
	GinEngine := gin.Default()
	GinEngine.Run()
}
