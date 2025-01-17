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

	server.GET("/users", users)

	// middleware (authenticated)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

}
