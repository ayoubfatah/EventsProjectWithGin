package main

import (
	"gin-quickstart/db"
	"gin-quickstart/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()


	config := cors.Config{
    	AllowOrigins:     []string{"http://localhost:3000", "https://eventify-fawn.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	server.Use(cors.New(config))


	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Backend is running!"})
	})


	routes.RegisterRoutes(server)

	// Get PORT from env or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}