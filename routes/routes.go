package routes

import (
	"github.com/gin-gonic/gin"
	"golang.com/rest/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/event/:id", getEvent)
	server.GET("/events", getEvents)

	server.POST("/signup", signup)
	server.POST("/login", login)

	// middleware (authenticated)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

}
