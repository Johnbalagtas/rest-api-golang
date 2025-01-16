package main

import (
	"github.com/gin-gonic/gin"
	"golang.com/rest/db"
	"golang.com/rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
