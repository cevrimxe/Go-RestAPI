package main

import (
	"github.com/cevrimxe/Go-RestAPI/db"
	"github.com/cevrimxe/Go-RestAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
