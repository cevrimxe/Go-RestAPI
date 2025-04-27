package routes

import (
	"github.com/cevrimxe/Go-RestAPI/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // get event by id

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)       // auth reqired
	authenticated.PUT("/events/:id", updateEvent)    // auth reqired and just update by who create it
	authenticated.DELETE("/events/:id", deleteEvent) // auth reqired and just update by who create it
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
