package main

import (
	"github.com/Kelvedler/GoBlog/models"
	"github.com/Kelvedler/GoBlog/server"
)

func main() {
	models.Init()
	defer models.CloseDBConn()
	server.Init()
}
