package routes

import "github.com/gin-gonic/gin"

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.GET("/users", getAllUsers)

	server.POST("/events", createEvents)
	server.POST("signup", signup)

	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
