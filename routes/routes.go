package routes

import (
	"gin-quickstart/middlewares"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine){
	// EVENTS
	server.GET("/events", getEvents)
	server.GET("/events/:slug", getEventBySlug)
	server.GET("/events/city/:city", getEventsByCity)
	server.GET("/events/me", getCurrentUserEvents)
	// 
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticated)
	server.POST("/events", createEvents )
	server.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// registration 
	authenticated.POST("/registration/:id", registerEvent)
	authenticated.DELETE("/registration/:id", cancelRegistration)


	// USERs
	server.POST("/signup", createUser)
	server.GET("/users", getAllUsers)
	authenticated.GET("/users/me", getCurrentUser) 
	server.POST("/login", login)
}
