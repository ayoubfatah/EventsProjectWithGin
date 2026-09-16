package routes

import (
	"gin-quickstart/middlewares"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine){
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticated)

	// EVENTS
	server.GET("/events", getEvents)
	server.GET("/events/:slug", getEventBySlug)
	server.GET("/events/city/:city", getEventsByCity)
	authenticated.GET("/events/registration", GetCurrentUserRegisteredEvents)



	authenticated.GET("/events/me", getCurrentUserEvents)
	authenticated.POST("/events", createEvents )
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// registration 
	authenticated.POST("/registration/:id", registerEvent)
	authenticated.DELETE("/registration/:id", cancelRegistration)
	authenticated.GET("/registration/:id", getEventReservation)


	// USERs
	server.POST("/signup", createUser)
	server.GET("/users", getAllUsers)
	authenticated.GET("/users/me", getCurrentUser) 
	server.POST("/login", login)
}
