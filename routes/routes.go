package routes

import (
	"gin-quickstart/middlewares"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine){
	// EVENTS
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	// 
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticated)
	authenticated.POST("/events", createEvents )
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// registration 
	authenticated.PUT("/registration/:id", registerEvent)
	authenticated.DELETE("/registration/:id", cancelRegistration)


	// USERs
	server.POST("/signup", createUser)
	server.GET("/users", getAllUsers)
	server.POST("/login", login)
}
